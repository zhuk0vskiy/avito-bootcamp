package dto

import "github.com/google/uuid"

type SubscribeRequest struct {
	UserID  uuid.UUID
	HouseID uuid.UUID
}

type SubscribeResponse struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	HouseID uuid.UUID
}

type UnsubscribeRequest struct {
	UserID  uuid.UUID
	HouseID uuid.UUID
}

type UnsubscribeResponse struct {
	ID      uuid.UUID
}
