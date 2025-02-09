package dto

import (
	"backend/internal/model"
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
	SubscriberID uuid.UUID
	ApartmentID  uuid.UUID
	HouseID      uuid.UUID
}

type AddNoticeResponse struct {
	NoticeID       uuid.UUID
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

type ProduceMessageRequest struct {
	Messages []*struct{
		CreationTime time.Time `json:"creation_time"`
		ApartmentID  uuid.UUID `json:"apartment_id"`
		SubscriberID uuid.UUID `json:"subscriber_id"`
		HouseID      uuid.UUID `json:"house_id"`
	}

}

type ProduceMessageResponse struct {}

type GetNoticesOutboxRequest struct {}

type GetNoticesOutboxResponse struct {
	Notices []*model.Notice
}

type ConfirmNoticeOutboxRequest struct {
	NoticeID uuid.UUID
}

type ConfirmNoticeOutboxResponse struct {
	NoticeOutboxID uuid.UUID
}