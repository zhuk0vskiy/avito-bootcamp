package dto

import (
	"time"

	"github.com/google/uuid"
)

type AddUserRequest struct {
	CreationTime time.Time
	Email        string
	Password     []byte
	IsModerator  bool
	TotpSecret   []byte
}

type AddUserResponse struct {
	ID uuid.UUID
	// CreationTime time.Time
}

type GetUserByEmailRequest struct {
	Email string
}

type GetUserByEmailResponse struct {
	ID           uuid.UUID
	CreationTime time.Time
	Email        string
	Password     []byte
	TotpSecret   []byte
	IsModerator  bool
}
