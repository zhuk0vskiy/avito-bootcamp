package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	CreateTime  time.Time
	Mail        string
	Password    string
	IsModerator bool
}

var (
	ErrUser_BadType     = errors.New("bd user type")
	ErrUser_BadRequest  = errors.New("bad nil request")
	ErrUser_BadMail     = errors.New("bad mail")
	ErrUser_BadPassword = errors.New("bad password")
)
