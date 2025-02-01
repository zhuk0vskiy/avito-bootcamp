package user

import (
	"backend/internal/repo/dto"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=UserIntf
type UserIntf interface {
	Add(ctx context.Context, request *dto.AddUserRequest) (*dto.AddUserResponse, error)
	GetByEmail(ctx context.Context, request *dto.GetUserByEmailRequest) (*dto.GetUserResponse, error)
}
