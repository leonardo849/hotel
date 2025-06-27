package repository

import (
	"fmt"
	"hotel/internal/dto"
	"hotel/internal/model"
)

type ReservationRepository struct {
	GuestRepository *GuestRepository
	RoomRepository *RoomRepository
}

func (r *ReservationRepository) CreateReservation(input dto.CreateReservationDTO)  error {
	if _, err := r.GuestRepository.FindOneGuest(input.GuestID.String()); err != nil {
		return  fmt.Errorf("that guest doesn't exist")
	}
	var roomDTO *dto.FindRoomDTO
	var err error
	if roomDTO, err = r.RoomRepository.FindOneRoom(input.RoomID.String()); err != nil {
		return fmt.Errorf("that room doesn't exist")
	}
	duration := input.CheckOut.Sub(input.CheckIn)
	days := int(duration.Hours() / 24)
	if days < 1 {
		return  fmt.Errorf("wtf bro")
	}
	price := roomDTO.PricePerNight * float64(days)
	reservation := model.Reservation{
		GuestID: input.GuestID,
		RoomID: input.RoomID,
		CheckIn: input.CheckIn,
		CheckOut: input.CheckOut,
		TotalPrice: price,
	}
	if err := DB.Create(&reservation).Error; err != nil {
		return  err
	}
	return  nil
	
}