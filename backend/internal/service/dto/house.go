package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateHouseRequest struct {
	CreatorID     uuid.UUID
	Address       string
	MaxApartments int32
}

type CreateHouseResponse struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	MaxApartments        int32
	UpdateApartmentsTime time.Time
}

type DeleteHouseRequest struct {
	ID uuid.UUID
}

// type DeleteHouseResponse struct {

// }

type GetHouseByIDRequest struct {
	ID uuid.UUID
}

type GetHouseByIDRespone struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Address              string
	MaxApartments        int32
	UpdateapartmentsTime time.Time
}
