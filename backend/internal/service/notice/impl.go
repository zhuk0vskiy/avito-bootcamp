package notice

import (
	repoDto "backend/internal/repo/dto"
	kafkaRepo "backend/internal/repo/kafka"
	noticeRepo "backend/internal/repo/postgres/notice"
	serviceDto "backend/internal/service/dto"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	// "golang.org/x/text/unicode/rangetable"
)

type NoticeService struct {
	logger         logger.Interface
	noticeRepoIntf noticeRepo.NoticeIntf
	kafkaRepoIntf  kafkaRepo.KafkaIntf
}

func NewNoticeService(logger logger.Interface, noticeRepoIntf noticeRepo.NoticeIntf, kafkaIntf kafkaRepo.KafkaIntf) *NoticeService {
	return &NoticeService{
		logger:         logger,
		noticeRepoIntf: noticeRepoIntf,
		kafkaRepoIntf:  kafkaIntf,
	}
}

func (s *NoticeService) Subscribe(ctx context.Context, request *serviceDto.SubscribeRequest) (*serviceDto.SubscribeResponse, error) {
	method := "NoticeService -- Create"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	err := validator.IsValidUUID(request.UserID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadUserID, err)
		return nil, ErrBadUserID
	}

	err = validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadHouseID, err)
		return nil, ErrBadHouseID
	}

	response, err := s.noticeRepoIntf.Subscribe(ctx, &repoDto.SubscribeRequest{
		CreationTime: time.Now(),
		SubsriberID:  request.UserID,
		HouseID:      request.HouseID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.SubscribeResponse{
		ID: response.ID,
	}, nil
}

// func (s *NoticeService) IsNeedToNoticeSubscribers(ctx context.Context, request *serviceDto.IsNeedToNoticeSubscribersRequest) (*serviceDto.IsNeedToNoticeSubscribersResponse, error) {
// 	method := "NoticeService -- IsNeedToNoticeSubscribers"
// 	s.logger.Infof("%s", method)
// 	if ctx == nil {
// 		s.logger.Errorf("%s -- %s", method, ErrNilContext)
// 		return nil, ErrNilContext
// 	}

// 	if request == nil {
// 		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
// 		return nil, ErrNilRequest
// 	}

// 	err := validator.IsValidUUID(request.HouseID.String())
// 	if err != nil {
// 		s.logger.Warnf("%s -- %s -- %s", method, ErrBadHouseID, err)
// 		return nil, ErrBadHouseID
// 	}

// 	response, err := s.noticeRepoIntf.GetSubscribersByHouseID(ctx, &repoDto.GetSubscribersByHouseIDRequest{
// 		HouseID: request.HouseID,
// 	})
// }

func (s *NoticeService) CreateNotices(ctx context.Context, request *serviceDto.CreateNoticesRequest) (*serviceDto.CreateNoticesResponse, error) {
	method := "NoticeService -- SendNotices"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	err := validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadHouseID, err)
		return nil, ErrBadHouseID
	}

	err = validator.IsValidUUID(request.ApartmentID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadApartmentID, err)
		return nil, ErrBadApartmentID
	}

	getResponse, err := s.noticeRepoIntf.GetSubscribersByHouseID(ctx, &repoDto.GetSubscribersByHouseIDRequest{
		HouseID: request.HouseID,
	})
	if err != nil {
		return nil, err
	}

	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var mtx sync.Mutex
	wg, ctx := errgroup.WithContext(ctxWithCancel)
	// var wg sync.WaitGroup

	noticeIDs := make([]uuid.UUID, 0)

	for _, userID := range getResponse.UsersIDs {
		wg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				addResponse, err := s.noticeRepoIntf.Add(ctx, &repoDto.AddNoticeRequest{
					CreationTime: time.Now(),
					SubscriberID: userID,
					ApartmentID:  request.ApartmentID,
					HouseID:      request.HouseID,
				})
				if err != nil {
					return err
				}

				mtx.Lock()
				noticeIDs = append(noticeIDs, addResponse.NoticeID)
				mtx.Unlock()
				return nil
			}
		})
	}

	err = wg.Wait()
	if err != nil {
		s.logger.Errorf("%s -- %s", method, ErrGoroutinesAdd)
	}

	return &serviceDto.CreateNoticesResponse{
		NoticeIDs: noticeIDs,
	}, nil

}

func (s *NoticeService) SendNoticesToKafka(ctx context.Context, request *serviceDto.SendNoticesToKafkaRequest) (*serviceDto.SendNoticesToKafkaResponse, error) {
	method := "NoticeService -- SendNoticesToKafka"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	getResponse, err := s.noticeRepoIntf.GetNoticesOutbox(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(getResponse)

	messages := make([]*struct {
		CreationTime time.Time "json:\"creation_time\""
		ApartmentID  uuid.UUID "json:\"apartment_id\""
		SubscriberID uuid.UUID "json:\"subscriber_id\""
		HouseID      uuid.UUID "json:\"house_id\""
	}, 0)

	var notices []uuid.UUID

	for _, notice := range getResponse.Notices {
		messages = append(messages, &struct {
			CreationTime time.Time "json:\"creation_time\""
			ApartmentID  uuid.UUID "json:\"apartment_id\""
			SubscriberID uuid.UUID "json:\"subscriber_id\""
			HouseID      uuid.UUID "json:\"house_id\""
		}{
			notice.CreationTime,
			notice.ApartmentID,
			notice.SubscriberID,
			notice.HouseID,
		})

		fmt.Println(notice)
		notices = append(notices, notice.ID)
	}

	repoRequest := &repoDto.ProduceMessageRequest{}
	repoRequest.Messages = messages

	_, err = s.kafkaRepoIntf.ProduceMessages(ctx, repoRequest)
	if err != nil {
		return nil, err
	}

	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctxWithCancel)

	for _, notice := range notices {
		wg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				_, err = s.noticeRepoIntf.ConfirmNoticeOutbox(ctx, &repoDto.ConfirmNoticeOutboxRequest{
					NoticeID: notice,
				})
				if err != nil {
					return err
				}
				return nil
			}
		})
	}

	err = wg.Wait()
	if err != nil {
		s.logger.Errorf("%s -- %s", method, ErrGoroutinesAdd)
	}

	return &serviceDto.SendNoticesToKafkaResponse{
		NoticeIDs: notices,
	}, nil
}
