package repository

import (
	"fmt"
	"hotel/internal/dto"
	"hotel/internal/model"

	"github.com/thoas/go-funk"
)

type RoomRepository struct {
	
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		
	}
}

func (r *RoomRepository) CreateRoom(input dto.CreateRoomDTO) error {
	room := model.Room{
		Number:        input.Number,
		Type:          input.Type,
		PricePerNight: input.PricePerNight,
	}
	return DB.Create(&room).Error
}

func (r *RoomRepository) FindAllRooms() ([]dto.FindRoomDTO, error) {
	var rooms []model.Room
	var mapped []dto.FindRoomDTO

	if err := DB.Find(&rooms).Preload("Reservations").Error; err != nil {
		return nil, err
	}

	mapped = funk.Map(rooms, func(room model.Room) dto.FindRoomDTO {
		reservations := room.Reservations
		mappedReservations := funk.Map(reservations, func(reservation model.Reservation) dto.FindReservationDTO {
			return dto.FindReservationDTO{
				ID:         reservation.ID,
				GuestID:    reservation.GuestID,
				Guest:      nil,
				CheckIn:    reservation.CheckIn,
				CheckOut:   reservation.CheckOut,
				TotalPrice: reservation.TotalPrice,
			}
		}).([]dto.FindReservationDTO)

		return dto.FindRoomDTO{
			ID:            room.ID,
			Number:        room.Number,
			Type:          room.Type,
			PricePerNight: room.PricePerNight,
			Reservations:  mappedReservations,
		}
	}).([]dto.FindRoomDTO)

	return mapped, nil
}

func (r *RoomRepository) FindOneRoom(id string) (*dto.FindRoomDTO, error) {
	var room model.Room
	if err := DB.First(&room, "id = ?", id).Preload("Reservations").Preload("Reservations.Guest").Error; err != nil {
		return nil, fmt.Errorf("room wasn't found")
	}

	mappedReservations := funk.Map(room.Reservations, func(reservation model.Reservation) dto.FindReservationDTO {
		guest := dto.FindGuestInReservationDTO{
			ID:        reservation.GuestID,
			Name:      reservation.Guest.Name,
			Email:     reservation.Guest.Email,
			Phone:     reservation.Guest.Phone,
			CreatedAt: reservation.CreatedAt,
			UpdatedAt: reservation.UpdatedAt,
		}
		return dto.FindReservationDTO{
			ID:         reservation.ID,
			GuestID:    reservation.GuestID,
			Guest:      &guest,
			CheckIn:    reservation.CheckIn,
			CheckOut:   reservation.CheckOut,
			TotalPrice: reservation.TotalPrice,
		}
	}).([]dto.FindReservationDTO)

	roomDTO := dto.FindRoomDTO{
		ID:            room.ID,
		Number:        room.Number,
		Type:          room.Type,
		PricePerNight: room.PricePerNight,
		Reservations:  mappedReservations,
	}

	return &roomDTO, nil

}

func (r *RoomRepository) UpdateRoom(id string, input dto.UpdateRoomDTO) error {
	fields := make(map[string]interface{})
	_, err := r.FindOneRoom(id)
	if err != nil {
		return  err
	}
	if input.Number != nil {
		fields["Number"] = *input.Number
	}
	if input.PricePerNight != nil {
		fields["PricePerNight"] = *input.PricePerNight
	}
	if input.Type != nil {
		fields["Type"] = *input.Type
	}
	if err := DB.Where("id = ?", id).Updates(fields).Error; err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) DeleteRoom(id string) error {
	_, err := r.FindOneRoom(id)
	if err != nil {
		return  err
	}
	if err := DB.Delete(id).Error; err != nil {
		return err
	}
	return nil
}
