package dto

import (
	"time"

	"github.com/google/uuid"
)

type AddHouseRequest struct {
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	MaxApartments        int32
	ApartmentsUpdateTime time.Time
}

type AddHouseResponse struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	MaxApartments        int32
	ApartmentsUpdateTime time.Time
}

type DeleteHouseRequest struct {
	ID uuid.UUID
}

// type DeleteHouseResponse struct {

// }

type GetHouseByIDRequest struct {
	ID uuid.UUID
}

type GetHouseByIDResponse struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	MaxApartments        int32
	ApartmentsUpdateTime time.Time
}
