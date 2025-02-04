package house

import (
	"backend/internal/repo/dto"
	"context"
	"errors"
)

type HouseIntf interface {
	Add(ctx context.Context, request *dto.AddHouseRequest) (*dto.AddHouseResponse, error)
	GetByID(ctx context.Context, request *dto.GetHouseByIDRequest) (*dto.GetHouseByIDResponse, error)
}

var (
	ErrUser_BadType     = errors.New("bd user type")
	ErrUser_BadMail     = errors.New("bad mail")
	ErrUser_BadPassword = errors.New("bad password")
	ErrUser_BadToken    = errors.New("bad token")
)

var (
	ErrExec       = errors.New("error while exec")
	ErrQueringRow = errors.New("error while quering row")
)

var (
	ErrNilRequest = errors.New("user add dto request is nil")
	ErrNilContext = errors.New("user add context is nil")
)
