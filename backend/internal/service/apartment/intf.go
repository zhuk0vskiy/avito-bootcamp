package apartment

import (
	"backend/internal/service/dto"
	"context"
	"errors"
)

type ApartmentIntf interface {
	Create(ctx context.Context, request *dto.CreateapartmentRequest) (*dto.CreateapartmentResponse, error)
	Delete(ctx context.Context, request *dto.CreateapartmentResponse) error
	GetByID(ctx context.Context, request *dto.GetByIDRequest) (*dto.GetByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateApartmentStatusRequest) (*dto.UpdateApartmentStatusResponse, error)
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
	ErrNilRequest = errors.New("dto request is nil")
	ErrNilContext = errors.New("context is nil")
)
