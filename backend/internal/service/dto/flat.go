package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateFlatRequest struct {
	CreatorID uuid.UUID
	HouseID   uuid.UUID
	Price     int
	Rooms     int
}

type CreateFlatResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int
	Rooms            int
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}

type DeleteFlatRequest struct {
	ID uuid.UUID
}

// type DeleteFlatResponse struct {

// }

type GetByIDRequest struct {
	ID uuid.UUID
}

type GetByIDResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int
	Rooms            int
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}

type GetByHouseIDRequest struct {
	HouseID uuid.UUID
}

type GetByHouseIDResponse struct {
	Flats []*struct {
		ID               uuid.UUID
		CreationTime     time.Time
		CreatorID        uuid.UUID
		HouseID          uuid.UUID
		Price            int
		Rooms            int
		Status           string
		StatusUpdateTime time.Time
		ModeratorID      uuid.UUID
	}
}

type UpdateFlatStatusRequest struct {
	ID          uuid.UUID
	Status      string
	ModeratorID uuid.UUID
}

type UpdateFlatStatusResponse struct {
	ID               uuid.UUID
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int
	Rooms            int
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
}
