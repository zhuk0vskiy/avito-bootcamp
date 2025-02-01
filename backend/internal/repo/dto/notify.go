package dto

import (
	"time"

	"github.com/google/uuid"
)

type SubscribeRequest struct {
	CreationTime time.Time
	SubsriberID  uuid.UUID
	HouseID      uuid.UUID
}

type SubscribeResponse struct {
	ID      uuid.UUID
	CreationTime time.Time
	SubscriberID  uuid.UUID
	HouseID uuid.UUID
}

type UnsubscribeRequest struct {
	SubscriberID  uuid.UUID
	HouseID uuid.UUID
}

type UnsubscribeResponse struct {
	ID uuid.UUID
}
