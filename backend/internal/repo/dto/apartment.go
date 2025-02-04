package dto

import (
	"backend/internal/model"
	"time"

	"github.com/google/uuid"
)

type AddApartmentRequest struct {
	CreationTime     time.Time
	CreatorID        uuid.UUID
	HouseID          uuid.UUID
	Price            int64
	Rooms            int32
	Status           string
	StatusUpdateTime time.Time
}

type AddApartmentResponse struct {
	ID uuid.UUID
	// CreationTime     time.Time
	// CreatorID        uuid.UUID
	// HouseID          uuid.UUID
	// Price            int64
	// Rooms            int32
	// Status           string
	// StatusUpdateTime time.Time
	// ModeratorID      uuid.UUID
}

type DeleteapartmentRequest struct {
	ID uuid.UUID
}

// type DeleteapartmentResponse struct {

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
	ID               uuid.UUID
	Status           string
	StatusUpdateTime time.Time
	ModeratorID      uuid.UUID
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
