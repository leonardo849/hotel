package integration_test

import (
	"hotel/internal/model"
	"testing"
)

func CreateReservationTest(t *testing.T) {
	e := newExpect(t)
	guest := model.Guest{
		Name: "batman2130930910",
		Email: "batman2@gmail.com",
		Phone: "333-332-9404",
	}
	if err := DB.Create(&model.Guest{}).Error; err != nil {
		t.Errorf("error creating guest")
	}
	room := model.Room{
		Number: 444,
		Type: "suite",
		PricePerNight: 300,
	}
	if err := DB.Create(&room).Error; err != nil {
		t.Errorf("error creating room")
	}

	

	e = newExpect(t)
	e.POST("/reservation/create"). 
	WithJSON(map[string]string{
		"room_id": room.ID.String(),
		"guest_id": guest.ID.String(),
		"check_in": "2025-06-01T12:00:00Z",
		"check_out": "2025-06-05T12:00:00Z",
	}). 
	Expect(). 
	Status(201)

}