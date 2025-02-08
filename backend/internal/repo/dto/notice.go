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
	ID           uuid.UUID
	CreationTime time.Time
	SubscriberID uuid.UUID
	HouseID      uuid.UUID
}

type UnsubscribeRequest struct {
	SubscriberID uuid.UUID
	HouseID      uuid.UUID
}

type UnsubscribeResponse struct {
	ID uuid.UUID
}

type AddNoticeRequest struct {
	CreationTime time.Time
	SubscriberID  uuid.UUID
	ApartmentID  uuid.UUID
	HouseID      uuid.UUID
}

type AddNoticeResponse struct {
	NoticeID uuid.UUID
	NoticeOutboxID uuid.UUID
}

// type SendNoticeRequest struct {
// 	ID           uuid.UUID
// 	CreationTime time.Time
// 	ApartmentID  uuid.UUID
// 	SubscriberID uuid.UUID
// 	HouseID      uuid.UUID
// }

type GetSubscribersByHouseIDRequest struct {
	HouseID uuid.UUID
}

type GetSubscribersByHouseIDResponse struct {
	UsersIDs []uuid.UUID
}