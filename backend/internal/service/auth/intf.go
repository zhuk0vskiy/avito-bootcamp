package auth

import (
	"backend/internal/service/dto"
	"context"

	"go.uber.org/zap"
)

type UserIntf interface {
	SignUp(ctx context.Context, logger *zap.Logger, request *dto.SignUpRequest) (*dto.SignUpResponse, error)
	LogIn(ctx context.Context, logger *zap.Logger, request *dto.LogInRequest) (*dto.LogInResponse, error)
}
