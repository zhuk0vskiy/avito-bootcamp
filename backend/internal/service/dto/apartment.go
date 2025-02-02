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
	Apartments []*struct {
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

type UpdateApartmentStatusRequest struct {
	ID          uuid.UUID
	Status      string
	ModeratorID uuid.UUID
}

type UpdateApartmentStatusResponse struct {
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
