package user

import (
	"backend/internal/model/dto"
	"context"
)

type UserIntr interface {
	SignUp(ctx context.Context, request *dto.SignUpRequest) error
	LogIn(ctx context.Context, request *dto.LogInRequest) (*dto.LogInResponse, error)
}