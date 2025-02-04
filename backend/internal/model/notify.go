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
	CreationTime time.Time
	ApartmentID  uuid.UUID
	HouseID      uuid.UUID
	SubscriberID uuid.UUID
	Status       string
}
