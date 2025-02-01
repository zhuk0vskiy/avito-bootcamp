package dto

import "time"

type AddUserRequest struct {
	CreationTime time.Time
	Email        string
	Password     []byte
	IsModerator  bool
	TotpSecret   []byte
}

type AddUserResponse struct {
	CreationTime time.Time
}

type GetUserRequest struct {
	Email string
}

type GetUserResponse struct {
	Email      string
	Password   []byte
	TotpSecret []byte
}
