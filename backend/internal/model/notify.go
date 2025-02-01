package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	SendedNotifyStatus   = "send"
	NoSendedNotifyStatus = "no send"
)

type Notify struct {
	ID           uuid.UUID
	CreateTime   time.Time
	FlatID       uuid.UUID
	HouseID      uuid.UUID
	SubscriberID uuid.UUID
	Status       string
}
