package house

import (
	"backend/internal/service/dto"
	"context"
	"errors"
)

type HouseIntf interface {
	Create(ctx context.Context, request *dto.CreateHouseRequest) (*dto.CreateHouseResponse, error)
	Delete(ctx context.Context, request *dto.DeleteHouseRequest) error
	GetByID(ctx context.Context, request *dto.GetByHouseIDRequest) (*dto.GetByHouseIDResponse, error)
}

var (
	ErrBadRequest        = errors.New("bad house request for create")
	ErrBadID             = errors.New("bad house id")
	ErrBadAddress        = errors.New("bad house address")
	ErrBadMaxAppartments = errors.New("bad house max apartments")
)

var (
	ErrNilRequest = errors.New("dto request is nil")
	ErrNilContext = errors.New("context is nil")
)
