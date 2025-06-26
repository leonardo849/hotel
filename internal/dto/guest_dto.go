package dto

import (
	"time"

	"github.com/google/uuid"
)

type FindGuestInReservationDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FindGuestDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Reservations []FindReservationDTO `json:"reservations"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateGuestDTO struct {
	Name  string `json:"name" validate:"required,min=10,max=200"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,phone_number"`
}
