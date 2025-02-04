package model

import (
	"time"

	"github.com/google/uuid"
)

type House struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	Maxapartments        int
	ApartmentsUpdateTime time.Time
}
