package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateHouseRequest struct {
	CreatorID     uuid.UUID
	Adress        string
	Maxapartments int
}

type CreateHouseResponse struct {
	ID                   uuid.UUID
	CreationTime         time.Time
	CreatorID            uuid.UUID
	Adress               string
	Maxapartments        int
	UpdateapartmentsTime time.Time
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
	Adress               string
	Maxapartments        int
	UpdateapartmentsTime time.Time
}
