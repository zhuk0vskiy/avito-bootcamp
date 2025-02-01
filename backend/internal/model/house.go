package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type House struct {
	ID              uuid.UUID
	CreateTime      time.Time
	CreatorID       uuid.UUID
	Adress          string
	MaxFlats        int
	UpdateFlatsTime time.Time
}

var (
	ErrHouse_BadRequest = errors.New("bad house request for create")
	ErrHouse_BadID      = errors.New("bad house id")
	ErrHouse_BadYear    = errors.New("bad house construct year")
)
