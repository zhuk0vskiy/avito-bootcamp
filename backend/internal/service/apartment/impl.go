package apartment

import (
	"backend/internal/model"
	apartmentRepo "backend/internal/repo/apartment"
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

func (s *apartmentService) Create(ctx context.Context, request *serviceDto.CreateapartmentRequest) (*serviceDto.CreateapartmentResponse, error) {
	method := "ApartmentServie -- Create"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilContext)
		return nil, model.ErrApartment_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilRequest)
		return nil, model.ErrApartment_NilRequest
	}

	err := validator.IsValidUUID(request.CreatorID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadCreatorID, err)
		return nil, model.ErrApartment_BadCreatorID
	}

	err = validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadHouseID, err)
		return nil, model.ErrApartment_BadCreatorID
	}

	if request.Price < 0 {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadPrice, err)
		return nil, model.ErrApartment_BadPrice
	}

	if request.Rooms < 0 {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadRooms, err)
		return nil, model.ErrApartment_BadRooms
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

	return &serviceDto.CreateapartmentResponse{
		ID:               addResponse.ID,
		CreationTime:     addResponse.CreationTime,
		CreatorID:        addResponse.CreatorID,
		HouseID:          addResponse.HouseID,
		Price:            addResponse.Price,
		Rooms:            addResponse.Rooms,
		Status:           addResponse.Status,
		StatusUpdateTime: addResponse.StatusUpdateTime,
		ModeratorID:      addResponse.ModeratorID,
	}, nil

}

func (s *apartmentService) GetByID(ctx context.Context, request *serviceDto.GetByIDRequest) (*serviceDto.GetByIDResponse, error) {
	method := "ApartmentServie -- GetByID"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilContext)
		return nil, model.ErrApartment_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilRequest)
		return nil, model.ErrApartment_NilRequest
	}

	err := validator.IsValidUUID(request.ID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadID, err)
		return nil, model.ErrApartment_BadID
	}

	response, err := s.apartmentRepo.GetByID(ctx, &repoDto.GetByIDRequest{
		ID: request.ID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.GetByIDResponse{
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

func (s *apartmentService) GetByHouseID(ctx context.Context, request *serviceDto.GetByHouseIDRequest) (*serviceDto.GetByHouseIDResponse, error) {
	method := "ApartmentServie -- GetByHouseID"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilContext)
		return nil, model.ErrApartment_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilRequest)
		return nil, model.ErrApartment_NilRequest
	}

	err := validator.IsValidUUID(request.HouseID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadID, err)
		return nil, model.ErrApartment_BadID
	}

	response, err := s.apartmentRepo.GetByHouseID(ctx, &repoDto.GetByHouseIDRequest{
		HouseID: request.HouseID,
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.GetByHouseIDResponse{
		Apartments: response.Apartments,
	}, nil
}

func (s *apartmentService) UpdateStatus(ctx context.Context, request *serviceDto.UpdateApartmentStatusRequest) (*serviceDto.UpdateApartmentStatusResponse, error) {
	method := "ApartmentServie -- UpdateStatus"
	s.logger.Infof("%s", method)
	if ctx == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilContext)
		return nil, model.ErrApartment_NilContext
	}

	if request == nil {
		s.logger.Errorf("%s -- %s", method, model.ErrApartment_NilRequest)
		return nil, model.ErrApartment_NilRequest
	}

	err := validator.IsValidUUID(request.ID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadID, err)
		return nil, model.ErrApartment_BadID
	}

	err = validator.IsValidUUID(request.ModeratorID.String())
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, model.ErrApartment_BadModeratorID, err)
		return nil, model.ErrApartment_BadModeratorID
	}

	if !validator.IsValidApartmentStatus(request.Status) {
		s.logger.Warnf("%s -- %s", method, model.ErrApartment_BadStatus)
		return nil, model.ErrApartment_BadStatus
	}

	response, err := s.apartmentRepo.UpdateStatus(ctx, &repoDto.UpdateapartmentStatusRequest{
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
