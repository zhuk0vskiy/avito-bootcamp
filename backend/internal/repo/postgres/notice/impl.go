package notice

import (
	repoDto "backend/internal/repo/dto"
	"backend/internal/repo/postgres"
	"backend/pkg/logger"
	"context"
	"fmt"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoticeRepo struct {
	logger       logger.Interface
	dbConnector  *pgxpool.Pool
	retryAdapter postgres.RetryAdapterIntf
}

func NewNoticeRepo(logger logger.Interface, dbConnector *pgxpool.Pool, retryAdapter postgres.RetryAdapterIntf) *NoticeRepo {
	return &NoticeRepo{
		logger:       logger,
		dbConnector:  dbConnector,
		retryAdapter: retryAdapter,
	}
}

func (r *NoticeRepo) Subscribe(ctx context.Context, request *repoDto.SubscribeRequest) (*repoDto.SubscribeResponse, error) {
	method := "NoticeRepo -- Subscribe"

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "insert into subscribers(creation_time, user_id, house_id) values ($1, $2, $3) returning id"
	// query := "insert into subscribers(creation_time, user_id, house_id) values ($1, $2, $3) on conflict (user_id, house_id) do update set creation_time = CURRENT_TIMESTAMP returning id"
	// query := "insert into subscribers(creation_time, user_id, house_id) values ($1, $2, $3) returning id on conflict (user_id, house_id) do returning id "
	response := repoDto.SubscribeResponse{}
	rows := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.CreationTime,
		request.SubsriberID,
		request.HouseID,
	)
	defer rows.Close()
	
	err := rows.Scan(
		&response.ID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s -- %s", method, ErrQueryRow, err)
		return nil, ErrQueryRow
	}
	return &response, nil
}

func (r *NoticeRepo) Add(ctx context.Context, request *repoDto.AddNoticeRequest) (*repoDto.AddNoticeResponse, error) {
	method := "NoticeRepo -- Add"

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	tx, err := r.dbConnector.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		r.logger.Warnf("%s -- %w -- %w", method, ErrStartTrans, err)
		return nil, ErrStartTrans
	}
	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			r.logger.Warnf("%s -- %w", method, ErrWhileTrns)
			if rollbackErr != nil {
				err = fmt.Errorf("%v -- %v", ErrRollback, err.Error())
			}
		}
	}()

	// response := repoDto.AddNoticeResponse{}
	var noticeID uuid.UUID
	var noticeOutbouxID uuid.UUID

	query := "insert into notices(creation_time, subscriber_id, apartment_id, house_id) values ($1, $2, $3, $4) returning id"
	err = tx.QueryRow(
		ctx,
		query,
		request.CreationTime,
		request.SubscriberID,
		request.ApartmentID,
		request.HouseID,
	).Scan(
		&noticeID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %w", method, ErrAddNotice)
		return nil, ErrAddNotice
	}

	query = "insert into notices_outbox(notice_id) values ($1) returning id"
	_, err = tx.Exec(
		ctx,
		query,
		&noticeOutbouxID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %w", method, ErrAddNoticeOutbox)
		return nil, ErrAddNoticeOutbox
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Warnf("%s -- %w", method, ErrCommit)
		return nil, ErrCommit
	}

	return &repoDto.AddNoticeResponse{
		NoticeID: noticeID,
		NoticeOutboxID: noticeOutbouxID,
	},nil
}

func (r *NoticeRepo) GetSubscribersByHouseID(ctx context.Context, request *repoDto.GetSubscribersByHouseIDRequest) (*repoDto.GetSubscribersByHouseIDResponse, error) {
	method := "NoticeRepo -- GetSubscribersByHouseID"

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "select user_id from subsribers where house_id = %1"

	rows, err := r.retryAdapter.Query(
		ctx,
		query,
		request.HouseID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQuery, err)
		return nil, ErrQuery
	}
	defer rows.Close()


	userIDs := make([]uuid.UUID, 0)
	var userID uuid.UUID
	// response := repoDto.GetApartmentsByHouseIDResponse{}
	// apartment := model.Apartment{}

	for rows.Next() {
		err = rows.Scan(
			&userID,
		)
		if err != nil {
			r.logger.Warnf("%s -- %s", method, ErrQueryRow, err)
			return nil, ErrQueryRow
		}
		userIDs = append(userIDs, userID)
	}

	return &repoDto.GetSubscribersByHouseIDResponse{
		UsersIDs: userIDs,
	}, nil

}
