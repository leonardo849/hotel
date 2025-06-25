package dto

import (
	"time"

	"github.com/google/uuid"
)

type FindReservationDTO struct {
	GuestID    uuid.UUID    `json:"guest_id"`
	Guest      *FindGuestDTO `json:"guest"`
	CheckIn    time.Time    `json:"check_in"`
	CheckOut   time.Time    `json:"check_out"`
	TotalPrice float64      `json:"total_price"`
}
