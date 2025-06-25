package dto

import (
	"hotel/internal/model"
	"time"
)

type FindGuestDTO struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Reservations *[]model.Reservation `json:"reservations"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}