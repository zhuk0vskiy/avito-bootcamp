package apartment

import (
	apartmentRepo "backend/internal/repo/postgres/apartment"
	repoDto "backend/internal/repo/dto"
	serviceDto "backend/internal/service/dto"
	"time"

	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"
)

type apartmentService struct {
	logger        logger.Interface
	apartmentRepo apartmentRepo.ApartmentIntf
}

func NewapartmentService(logger logger.Interface, flarRepo apartmentRepo.ApartmentIntf) *apartmentService {
	return &apartmentService{
		logger:        logger,
		apartmentRepo: flarRepo,
	}
}

func (s *apartmentService) Create(ctx context.Context, request *serviceDto.CreateApartmentRequest) (*serviceDto.CreateApartmentResponse, error) {
	method := "ApartmentService -- Create"
	// s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	err := validator.IsValidUUID(request.CreatorID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s", method, err)
		return nil, ErrBadCreatorID
	}

	err = validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s", method, err)
		return nil, ErrBadCreatorID
	}

	if request.Price < 0 {
		s.logger.Warnf("%s -- %s", method, ErrBadPrice)
		return nil, ErrBadPrice
	}

	if request.Rooms < 0 {
		s.logger.Warnf("%s -- %s", method, ErrBadRooms)
		return nil, ErrBadRooms
	}

	addResponse, err := s.apartmentRepo.Add(ctx, &repoDto.AddApartmentRequest{
		CreationTime:     time.Now(),
		CreatorID:        request.CreatorID,
		HouseID:          request.HouseID,
		Price:            request.Price,
		Rooms:            request.Rooms,
		Status:           "created",
		StatusUpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.CreateApartmentResponse{
		ID:               addResponse.ID,
		// CreationTime:     addResponse.CreationTime,
		// CreatorID:        addResponse.CreatorID,
		// HouseID:          addResponse.HouseID,
		// Price:            addResponse.Price,
		// Rooms:            addResponse.Rooms,
		// Status:           addResponse.Status,
		// StatusUpdateTime: addResponse.StatusUpdateTime,
		// ModeratorID:      addResponse.ModeratorID,
	}, nil

}

func (s *apartmentService) GetByID(ctx context.Context, request *serviceDto.GetApartmentByIDRequest) (*serviceDto.GetApartmentByIDResponse, error) {
	method := "ApartmentServie -- GetByID"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	err := validator.IsValidUUID(request.ID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
	}

	response, err := s.apartmentRepo.GetByID(ctx, &repoDto.GetApartmentByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.GetApartmentByIDResponse{
		ID:               response.ID,
		CreationTime:     response.CreationTime,
		CreatorID:        response.CreatorID,
		HouseID:          response.HouseID,
		Price:            response.Price,
		Rooms:            response.Rooms,
		Status:           response.Status,
		StatusUpdateTime: response.StatusUpdateTime,
		ModeratorID:      response.ModeratorID,
	}, nil

}

func (s *apartmentService) GetByHouseID(ctx context.Context, request *serviceDto.GetApartmentsByHouseIDRequest) (*serviceDto.GetApartmentsByHouseIDResponse, error) {
	method := "ApartmentServie -- GetByHouseID"
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

	response, err := s.apartmentRepo.GetByHouseID(ctx, &repoDto.GetApartmentsByHouseIDRequest{
		HouseID: request.HouseID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.GetApartmentsByHouseIDResponse{
		Apartments: response.Apartments,
	}, nil
}

func (s *apartmentService) UpdateStatus(ctx context.Context, request *serviceDto.UpdateApartmentStatusRequest) (*serviceDto.UpdateApartmentStatusResponse, error) {
	method := "ApartmentServie -- UpdateStatus"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilContext)
		return nil, ErrNilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, ErrNilRequest)
		return nil, ErrNilRequest
	}

	err := validator.IsValidUUID(request.ID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
	}

	err = validator.IsValidUUID(request.ModeratorID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadModeratorID, err)
		return nil, ErrBadModeratorID
	}

	if !validator.IsValidApartmentStatus(request.Status) {
		s.logger.Warnf("%s -- %s", method, ErrBadStatus)
		return nil, ErrBadStatus
	}

	response, err := s.apartmentRepo.UpdateStatus(ctx, &repoDto.UpdateApartmentStatusRequest{
		ID:               request.ID,
		Status:           request.Status,
		StatusUpdateTime: time.Now(),
		ModeratorID:      request.ModeratorID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.UpdateApartmentStatusResponse{
		ID:               response.ID,
		CreationTime:     response.CreationTime,
		CreatorID:        response.CreatorID,
		HouseID:          response.HouseID,
		Price:            response.Price,
		Rooms:            response.Rooms,
		Status:           response.Status,
		StatusUpdateTime: response.StatusUpdateTime,
		ModeratorID:      response.ModeratorID,
	}, nil

}
