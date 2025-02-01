package apartment

import (
	"backend/internal/model"
	repoDto "backend/internal/repo/dto"
	serviceDto "backend/internal/service/dto"
	apartmentRepo "backend/internal/repo/apartment"
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
	method := "apartmentServie -- Create"
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
