package dto

import (
	"time"

	"github.com/google/uuid"
)

type SignUpRequest struct {
	Email    string
	Password string
	Role     string
}

type SignUpResponse struct {
	ID           uuid.UUID
	CreationTime time.Time
	TotpSecret   string
}

type LogInRequest struct {
	Email    string
	Password string
	Token    string
}

type LogInResponse struct {
	Token string
}
