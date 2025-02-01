package user

import (
	"backend/internal/repo/dto"
	"context"
)

type UserIntr interface {
	Add(ctx context.Context, request *dto.AddUserRequest) (*dto.AddUserResponse, error)
	GetByEmail(ctx context.Context, request *dto.GetUserRequest) (*dto.GetUserResponse, error)
}