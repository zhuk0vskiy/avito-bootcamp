package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Flat struct {
	ID               uuid.UUID
	CreateTime       time.Time
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
	AnyStatus        = "any"
)

var (
	ErrFlat_BadPrice   = errors.New("bad flat price")
	ErrFlat_BadID      = errors.New("bad flat id")
	ErrFlat_BadHouseID = errors.New("bad flats house id")
	ErrFlat_BadRooms   = errors.New("bad flat rooms")
	ErrFlat_BadNewFlat = errors.New("bad new flat for update")
	ErrFlat_BadStatus  = errors.New("bad flat status")
	ErrFlat_BadRequest = errors.New("bad request for create")
)
