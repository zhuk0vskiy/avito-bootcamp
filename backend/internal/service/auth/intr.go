package user

import (
	"backend/internal/service/dto"
	"context"
)

type UserIntr interface {
	SignUp(ctx context.Context, request *dto.SignUpRequest) error
	LogIn(ctx context.Context, request *dto.LogInRequest) (*dto.LogInResponse, error)
}