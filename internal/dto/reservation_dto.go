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
