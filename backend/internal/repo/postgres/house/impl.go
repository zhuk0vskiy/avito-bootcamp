package house

import (
	repoDto "backend/internal/repo/dto"
	"backend/internal/repo/postgres"
	"backend/pkg/logger"
	"context"
)

type HouseRepo struct {
	logger       logger.Interface
	retryAdapter postgres.RetryAdapterIntf
}

func NewHouseRepo(logger logger.Interface, retryAdapter postgres.RetryAdapterIntf) *HouseRepo {
	return &HouseRepo{
		logger:       logger,
		retryAdapter: retryAdapter,
	}
}

func (r *HouseRepo) Add(ctx context.Context, request *repoDto.AddHouseRequest) (*repoDto.AddHouseResponse, error) {
	method := "UserRepo -- Add"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}
	query := "insert into houses(creation_time, creator_id, address, max_apartments, apartments_update_time) values ($1, $2, $3, $4, $5) returning id"

	// query := "insert into houses(creation_time, creator_id, address, max_flats, last_update_time) values ($1, $2, $3, $4, $5) returning id"

	response := repoDto.AddHouseResponse{}
	rows := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.CreationTime,
		request.CreatorID,
		request.Address,
		request.MaxApartments,
		request.ApartmentsUpdateTime,
	)
	defer rows.Close()

	err := rows.Scan(
		&response.ID,
	)
	
	if err != nil {
		r.logger.Warnf("%s -- %s -- %s", method, ErrQueringRow, err)
		return nil, ErrQueringRow
	}

	return &response, nil
}

func (r *HouseRepo) GetByID(ctx context.Context, request *repoDto.GetHouseByIDRequest) (*repoDto.GetHouseByIDResponse, error) {
	method := "UserRepo -- Add"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "select id, creation_time, creator_id, address, max_apartments, update_apartments_time from houses where id = $1"

	response := repoDto.GetHouseByIDResponse{}
	err := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.ID,
	).Scan(
		response.ID,
		response.CreationTime,
		response.CreatorID,
		response.Address,
		response.MaxApartments,
		response.ApartmentsUpdateTime,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQueringRow, err)
		return nil, ErrQueringRow
	}

	return &response, nil
}
