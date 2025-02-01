package model

import (
	"errors"
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
	AnyStatus        = "any"
)

var (
	ErrApartment_BadPrice        = errors.New("bad apartment price")
	ErrApartment_BadID           = errors.New("bad apartment id")
	ErrApartment_BadHouseID      = errors.New("bad apartments house id")
	ErrApartment_BadCreatorID    = errors.New("bad apartments creator id")
	ErrApartment_BadRooms        = errors.New("bad apartment rooms")
	ErrApartment_BadNewapartment = errors.New("bad new apartment for update")
	ErrApartment_BadStatus       = errors.New("bad apartment status")
	ErrApartment_BadRequest      = errors.New("bad request for create")
)

var (
	ErrApartment_NilRequest = errors.New("dto request is nil")
	ErrApartment_NilContext = errors.New("context is nil")
)
