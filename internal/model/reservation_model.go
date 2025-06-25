package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	GuestID    uuid.UUID `json:"guest_id"`
	RoomID     uuid.UUID `json:"room_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	Guest      Guest     `gorm:"foreignKey:GuestID;references:ID"`
	Room       Room      `gorm:"foreignKey:RoomID;references:ID"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (r *Reservation) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
