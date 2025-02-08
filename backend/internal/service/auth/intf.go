package auth

import (
	"backend/internal/service/dto"
	"context"
	"errors"
)

type AuthIntf interface {
	SignUp(ctx context.Context, request *dto.SignUpRequest) (*dto.SignUpResponse, error)
	LogIn(ctx context.Context, request *dto.LogInRequest) (*dto.LogInResponse, error)
}

var (
	ErrBadType     = errors.New("bd user type")
	ErrBadMail     = errors.New("bad mail")
	ErrBadPassword = errors.New("bad password")
	ErrBadToken    = errors.New("bad token")
)

var (
	ErrErrorHashPassword = errors.New("error with generate passwotd hash")
	ErrErrorTotpGenerate = errors.New("error with generate totp")
	ErrTotpEncrypt       = errors.New("error with totp encrypt")
	ErrTotpDecrypt       = errors.New("error with totp decrypt")
)

var (
	ErrNilRequest = errors.New("dto request is nil")
	ErrNilContext = errors.New("context is nil")
)
