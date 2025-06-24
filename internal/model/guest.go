package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Guest struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Reservations []Reservation `gorm:"foreignKey:GuestID;references:ID"`
	CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

func (r *Guest) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return 
}