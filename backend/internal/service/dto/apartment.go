package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateapartmentRequest struct {
	CreatorID uuid.UUID
	HouseID   uuid.UUID
	Price     int64
	Rooms     int32
}

type CreateapartmentResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int64
	Rooms            int32
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}

type DeleteapartmentRequest struct {
	ID uuid.UUID
}

// type DeleteapartmentResponse struct {

// }

type GetByIDRequest struct {
	ID uuid.UUID
}

type GetByIDResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int64
	Rooms            int32
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}

type GetByHouseIDRequest struct {
	HouseID uuid.UUID
}

type GetByHouseIDResponse struct {
	apartments []*struct {
		ID               uuid.UUID
		CreationTime     time.Time
		CreatorID        uuid.UUID
		HouseID          uuid.UUID
		Price            int64
		Rooms            int32
		Status           string
		StatusUpdateTime time.Time
		ModeratorID      uuid.UUID
	}
}

type UpdateapartmentStatusRequest struct {
	ID          uuid.UUID
	Status      string
	ModeratorID uuid.UUID
}

type UpdateapartmentStatusResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int64
	Rooms            int32
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}
