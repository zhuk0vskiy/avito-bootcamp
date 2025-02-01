package dto

import (
	"time"

	"github.com/google/uuid"
)

type AddHouseRequest struct {
	CreationTime    time.Time
	CreatorID       uuid.UUID
	Adress          string
	MaxFlats        int
	UpdateFlatsTime time.Time
}

type AddHouseResponse struct {
	ID              uuid.UUID
	CreationTime    time.Time
	CreatorID       uuid.UUID
	Adress          string
	MaxFlats        int
	UpdateFlatsTime time.Time
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
	ID              uuid.UUID
	CreationTime    time.Time
	CreatorID       uuid.UUID
	Adress          string
	MaxFlats        int
	UpdateFlatsTime time.Time
}
