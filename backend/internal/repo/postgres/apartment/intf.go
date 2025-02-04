package apartment

import (
	"backend/internal/repo/dto"
	"context"
	"errors"
)

type ApartmentIntf interface {
	Add(ctx context.Context, request *dto.AddApartmentRequest) (*dto.AddApartmentResponse, error)
	GetByID(ctx context.Context, request *dto.GetApartmentByIDRequest) (*dto.GetApartmentByIDResponse, error)
	GetByHouseID(ctx context.Context, request *dto.GetApartmentsByHouseIDRequest) (*dto.GetApartmentsByHouseIDResponse, error)
	UpdateStatus(ctx context.Context, request *dto.UpdateApartmentStatusRequest) (*dto.UpdateApartmentStatusResponse, error)
}

var (
	ErrUser_BadType     = errors.New("bd user type")
	ErrUser_BadMail     = errors.New("bad mail")
	ErrUser_BadPassword = errors.New("bad password")
	ErrUser_BadToken    = errors.New("bad token")
)

var (
	ErrExec       = errors.New("error while exec")
	ErrQueryRow = errors.New("error while quering row")
	ErrQuery      = errors.New("error while query")
)

var (
	ErrNilRequest = errors.New("user add dto request is nil")
	ErrNilContext = errors.New("user add context is nil")
)
