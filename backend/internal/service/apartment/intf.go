package apartment

import (
	serviceDto "backend/internal/service/dto"
	"context"
	"errors"
)

type ApartmentIntf interface {
	Create(ctx context.Context, request *serviceDto.CreateApartmentRequest) (*serviceDto.CreateApartmentResponse, error)
	// Delete(ctx context.Context, request *serviceDto.CreateApartmentResponse) error
	GetByID(ctx context.Context, request *serviceDto.GetApartmentByIDRequest) (*serviceDto.GetApartmentByIDResponse, error)
	GetByHouseID(ctx context.Context, request *serviceDto.GetApartmentsByHouseIDRequest) (*serviceDto.GetApartmentsByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *serviceDto.UpdateApartmentStatusRequest) (*serviceDto.UpdateApartmentStatusResponse, error)
}

var (
	ErrBadPrice        = errors.New("bad apartment price")
	ErrBadID           = errors.New("bad apartment id")
	ErrBadHouseID      = errors.New("bad apartments house id")
	ErrBadCreatorID    = errors.New("bad apartments creator id")
	ErrBadModeratorID  = errors.New("bad apartments moderator id")
	ErrBadRooms        = errors.New("bad apartment rooms")
	ErrBadNewapartment = errors.New("bad new apartment for update")
	ErrBadStatus       = errors.New("bad apartment status")
	ErrBadRequest      = errors.New("bad request for create")
)

var (
	ErrNilRequest = errors.New("serviceDto request is nil")
	ErrNilContext = errors.New("context is nil")
)
