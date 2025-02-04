package model

import (
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

