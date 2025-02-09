package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	SendedNoticeStatus   = "send"
	NoSendedNoticeStatus = "no send"
)

type Notice struct {
	ID           uuid.UUID `json:"id"`
	CreationTime time.Time `json:"creation_time"`
	ApartmentID  uuid.UUID `json:"apartment_id"`
	SubscriberID uuid.UUID `json:"subscriber_id"`
	HouseID      uuid.UUID `json:"house_id"`
}
