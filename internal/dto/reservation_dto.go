package dto

import (
	"time"

	"github.com/google/uuid"
)

type FindReservationDTO struct {
	ID         uuid.UUID                  `json:"id"`
	RoomID     uuid.UUID                  `json:"room_id"`
	GuestID    uuid.UUID                  `json:"guest_id"`
	Guest      *FindGuestInReservationDTO `json:"guest"`
	CheckIn    time.Time                  `json:"check_in"`
	CheckOut   time.Time                  `json:"check_out"`
	TotalPrice float64                    `json:"total_price"`
}

type CreateReservationDTO struct {
    RoomID   uuid.UUID `json:"room_id" validate:"required"`
    GuestID  uuid.UUID `json:"guest_id" validate:"required"`
    CheckIn  time.Time `json:"check_in" validate:"required"`
    CheckOut time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
}
