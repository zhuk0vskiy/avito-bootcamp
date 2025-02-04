package dto

import (
	"backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type CreateApartmentRequest struct {
	CreatorID uuid.UUID
	HouseID   uuid.UUID
	Price     int64
	Rooms     int32
}

type CreateApartmentResponse struct {
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

type DeleteApartmentRequest struct {
	ID uuid.UUID
}

// type DeleteApartmentResponse struct {

// }

type GetApartmentByIDRequest struct {
	ID uuid.UUID
}

type GetApartmentByIDResponse struct {
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

type GetApartmentsByHouseIDRequest struct {
	HouseID uuid.UUID
}

type GetApartmentsByHouseIDResponse struct {
	Apartments []*model.Apartment
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
