package notice

import (
	repoDto "backend/internal/repo/dto"
	noticeRepo "backend/internal/repo/postgres/notice"
	serviceDto "backend/internal/service/dto"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"
	"sync"
	"time"
	// "golang.org/x/text/unicode/rangetable"
)

type NoticeService struct {
	mtx sync.Mutex
	logger         logger.Interface
	noticeRepoIntf noticeRepo.NoticeIntf
}

func NewNoticeService(logger logger.Interface, noticeRepoIntf noticeRepo.NoticeIntf) *NoticeService {
	return &NoticeService{
		logger:         logger,
		noticeRepoIntf: noticeRepoIntf,
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
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
	}

	err = validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
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

func (s *NoticeService) IsNeedToNoticeSubscribers(ctx context.Context, request *serviceDto.IsNeedToNoticeSubscribersRequest) (*serviceDto.IsNeedToNoticeSubscribersResponse, error) {
	method := "NoticeService -- IsNeedToNoticeSubscribers"
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
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
	}

	response, err := s.noticeRepoIntf.GetSubscribersByHouseID(ctx, &repoDto.GetSubscribersByHouseIDRequest{
		HouseID: request.HouseID,
	})
}

func (s *NoticeService) SendNotices(ctx context.Context, request *serviceDto.SendNoticesRequest) (*serviceDto.SendNoticesResponse, error) {
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

	response, err := s.noticeRepoIntf.GetSubscribersByHouseID(ctx, &repoDto.GetSubscribersByHouseIDRequest{
		HouseID: request.HouseID,
	})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup


	for _, userID := range(response.UsersIDs) {
		wg.Add(1)
		go func() {
			s.noticeRepoIntf.Add(ctx, &repoDto.AddNoticeRequest{
				CreationTime: time.Now(),
				SubscriberID: userID,
			})

			// if 
		}()
	}

	wg.Wait()
	


	// for _, userID := range(response.UsersIDs) {
	// 	err = validator.IsValidUUID(userID.String())
	// 	if err != nil {
	// 		s.logger.Warnf("%s -- %s -- %s", method, ErrBadApartmentID, err)
	// 		return nil, ErrBadApartmentID
	// 	}
	// }




}
 