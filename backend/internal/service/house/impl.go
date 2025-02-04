package house

import (
	repoDto "backend/internal/repo/dto"
	houseRepo "backend/internal/repo/postgres/house"
	serviceDto "backend/internal/service/dto"
	"backend/pkg/logger"
	"backend/pkg/validator"
	"context"
	"time"
)

type HouseService struct {
	logger        logger.Interface
	houseRepoIntf houseRepo.HouseIntf
}

func NewHouseService(logger logger.Interface, houseRepoIntf houseRepo.HouseIntf) *HouseService {
	return &HouseService{
		logger:        logger,
		houseRepoIntf: houseRepoIntf,
	}
}

func (s *HouseService) Create(ctx context.Context, request *serviceDto.CreateHouseRequest) (*serviceDto.CreateHouseResponse, error) {
	method := "HouseServie -- Create"
	s.logger.Infof("%s", method)
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
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadID, err)
		return nil, ErrBadID
	}

	err = validator.ValidateAddress(request.Address)
	if err != nil {
		s.logger.Warnf("%s -- %s -- %s", method, ErrBadAddress, err)
		return nil, ErrBadAddress
	}

	if request.MaxApartments <= 0 {
		s.logger.Warnf("%s -- %s", ErrBadMaxAppartments)
		return nil, ErrBadMaxAppartments
	}

	response, err := s.houseRepoIntf.Add(ctx, &repoDto.AddHouseRequest{
		CreationTime:         time.Now(),
		CreatorID:            request.CreatorID,
		Address:              request.Address,
		MaxApartments:        request.MaxApartments,
		ApartmentsUpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &serviceDto.CreateHouseResponse{
		ID:                   response.ID,
		CreationTime:         response.CreationTime,
		CreatorID:            response.CreatorID,
		Address:              response.Address,
		MaxApartments:        response.MaxApartments,
		ApartmentsUpdateTime: response.ApartmentsUpdateTime,
	}, nil
}
