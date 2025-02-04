package user

import (
	"backend/internal/repo/dto"
	"context"
	"errors"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=UserIntf
type UserIntf interface {
	Add(ctx context.Context, request *dto.AddUserRequest) (*dto.AddUserResponse, error)
	GetByEmail(ctx context.Context, request *dto.GetUserByEmailRequest) (*dto.GetUserResponse, error)
}



var (
	ErrUser_BadType     = errors.New("bd user type")
	ErrUser_BadMail     = errors.New("bad mail")
	ErrUser_BadPassword = errors.New("bad password")
	ErrUser_BadToken    = errors.New("bad token")
)

var (
	ErrExec = errors.New("error while exec")
	ErrQueringRow = errors.New("error while quering row")
)

var (
	ErrNilRequest = errors.New("dto request is nil")
	ErrNilContext = errors.New("context is nil")
)
