package model

import (
	"time"

	"github.com/google/uuid"
)

type Apartment struct {
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

const (
	CreatedStatus    = "created"
	ApprovedStatus   = "approved"
	DeclinedStatus   = "declined"
	ModeratingStatus = "on moderation"
	// AnyStatus        = "any"
)