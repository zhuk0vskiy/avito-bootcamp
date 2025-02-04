package apartment

import (
	"backend/internal/model"
	repoDto "backend/internal/repo/dto"
	"backend/internal/repo/postgres"
	"backend/pkg/logger"
	"context"
)

type ApartmentRepo struct {
	logger       logger.Interface
	retryAdapter postgres.RetryAdapterIntf
}

func NewApartmentRepo(logger logger.Interface, retryAdapter postgres.RetryAdapterIntf) *ApartmentRepo {
	return &ApartmentRepo{
		logger:       logger,
		retryAdapter: retryAdapter,
	}
}

func (r *ApartmentRepo) Add(ctx context.Context, request *repoDto.AddApartmentRequest) (*repoDto.AddApartmentResponse, error) {
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

	query := "insert into apartments(creation_time, creator_id, house_id, price, rooms, status, status_update_time) values ($1, $2, $3, $4, $5, $6, $7) returning id"

	response := repoDto.AddApartmentResponse{}
	err := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.CreationTime,
		request.CreatorID,
		request.HouseID,
		request.Price,
		request.Rooms,
		request.Status,
		request.StatusUpdateTime,
	).Scan(
		response.ID,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQueryRow, err)
		return nil, ErrQueryRow
	}
	return &response, nil
}

func (r *ApartmentRepo) GetByID(ctx context.Context, request *repoDto.GetApartmentByIDRequest) (*repoDto.GetApartmentByIDResponse, error) {
	method := "UserRepo -- GetByID"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "select id, creation_time, creator_id, house_id, price, rooms, status, status_update_time from apartments where id = $1"

	response := repoDto.GetApartmentByIDResponse{}
	err := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.ID,
	).Scan(
		response.ID,
		response.CreationTime,
		response.CreatorID,
		response.HouseID,
		response.Price,
		response.Rooms,
		response.Status,
		response.StatusUpdateTime,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQueryRow, err)
		return nil, ErrQueryRow
	}

	return &response, nil
}

func (r *ApartmentRepo) GetByHouseID(ctx context.Context, request *repoDto.GetApartmentsByHouseIDRequest) (*repoDto.GetApartmentsByHouseIDResponse, error) {
	method := "UserRepo -- GetByHouseID"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "select id, creation_time, creator_id, house_id, price, rooms, status, status_update_time from apartments where house_id = $1"

	rows, err := r.retryAdapter.Query(
		ctx,
		query,
		request.HouseID,
	)
	defer rows.Close()
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQuery, err)
		return nil, ErrQuery
	}

	response := repoDto.GetApartmentsByHouseIDResponse{}
	apartment := model.Apartment{}

	for rows.Next() {
		err = rows.Scan(
			&apartment.ID,
			&apartment.CreationTime,
			&apartment.CreatorID,
			&apartment.HouseID,
			&apartment.Price,
			&apartment.Rooms,
			&apartment.Status,
			&apartment.StatusUpdateTime,
		)
		if err != nil {
			r.logger.Warnf("%s -- %s", method, ErrQueryRow, err)
			return nil, ErrQueryRow
		}
		response.Apartments = append(response.Apartments, &apartment)
	}

	return &response, nil
}

func (r *ApartmentRepo) UpdateStatus(ctx context.Context, request *repoDto.UpdateApartmentStatusRequest) (*repoDto.UpdateApartmentStatusResponse, error) {
	method := "UserRepo -- GetByHouseID"
	r.logger.Infof("%s", method)

	if ctx == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		r.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	query := "update houses set status = $1, status_update_time = $2, moderator_id = $3 where id = $4 returning iid, creation_time, creator_id, house_id, price, rooms, status, status_update_time"

	response := repoDto.UpdateApartmentStatusResponse{}
	err := r.retryAdapter.QueryRow(
		ctx,
		query,
		request.Status,
		request.StatusUpdateTime,
		request.ModeratorID,
		request.ID,
	).Scan(
		response.ID,
		response.CreationTime,
		response.CreatorID,
		response.HouseID,
		response.Price,
		response.Rooms,
		response.Status,
		response.StatusUpdateTime,
	)
	if err != nil {
		r.logger.Warnf("%s -- %s", method, ErrQueryRow, err)
		return nil, ErrQueryRow
	}

	return &response, nil

}

