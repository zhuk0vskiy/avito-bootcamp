package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	CreationTime time.Time
	Mail         string
	Password     string
	IsModerator  bool
}

var (
	ErrUser_BadType     = errors.New("bd user type")
	ErrUser_BadMail     = errors.New("bad mail")
	ErrUser_BadPassword = errors.New("bad password")
	ErrUser_BadToken = errors.New("bad token")
)

var (
	ErrUser_ErrorHashPassword = errors.New("error with generate passwotd hash")
	ErrUser_ErrorTotpGenerate = errors.New("error with generate totp")
	ErrUser_TotpEncrypt       = errors.New("error with totp encrypt")
	ErrUser_TotpDecrypt       = errors.New("error with totp decrypt")
)

var (
	ErrUser_NilRequest = errors.New("dto request is nil")
	ErrUser_NilContext = errors.New("context is nil")
)
