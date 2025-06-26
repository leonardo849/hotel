package dto

import "github.com/google/uuid"

type CreateRoomDTO struct {
	Number        uint    `validate:"required,number,gt=0" json:"number"`
	Type          string  `json:"type" validate:"required,typeofroom"`
	PricePerNight float64 `json:"price_per_night" validate:"required,gt=0,number"`
}

type UpdateRoomDTO struct {
	Number        *uint    `validate:"omitempty,number,gt=0" json:"number"`
	Type          *string  `json:"type" validate:"omitempty,typeofroom"`
	PricePerNight *float64 `json:"price_per_night" validate:"omitempty,gt=0"`
}

type FindRoomDTO struct {
	ID            uuid.UUID            `json:"id"`
	Number        uint                 `json:"number"`
	Type          string               `json:"type" `
	PricePerNight float64              `json:"price_per_night"`
	Reservations  []FindReservationDTO `json:"reservations"`
}