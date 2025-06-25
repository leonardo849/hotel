package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Number        uint      `json:"number"`
	Type          string    `json:"type"`
	PricePerNight float64   `json:"price_per_night"`
	Reservations []Reservation `gorm:"foreignKey:RoomID;references:ID" json:"reservations"` 
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
