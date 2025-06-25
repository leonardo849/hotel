package repository

import (
	"hotel/internal/dto"
	"hotel/internal/model"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type RoomRepository struct {
	roomModel *gorm.DB
}

func NewRoomRepository(roomModel *gorm.DB) *RoomRepository {
	return &RoomRepository{
		roomModel: roomModel,
	}
}

func (r *RoomRepository) CreateRoom(input dto.CreateRoomDTO) error {
	room := model.Room{
		Number: input.Number,
		Type: input.Type,
		PricePerNight: input.PricePerNight,
	}
	return r.roomModel.Create(&room).Error
}

func (r *RoomRepository) FindAllRooms() ([]dto.FindRoomDTO, error) {
	var rooms []model.Room
	var mapped []dto.FindRoomDTO 

	if err := r.roomModel.Find(&rooms).Preload("Reservations").Error; err != nil {
		return  nil, err
	}

	mapped = funk.Map(rooms, func(room model.Room) dto.FindRoomDTO {
		reservations := room.Reservations
		mappedReservations := funk.Map(reservations, func(reservation model.Reservation) dto.FindReservationDTO {
			return dto.FindReservationDTO{
				GuestID: reservation.GuestID,
				Guest: nil,
				CheckIn: reservation.CheckIn,
				CheckOut: reservation.CheckOut,
				TotalPrice: reservation.TotalPrice,
			}
		}).([]dto.FindReservationDTO)

		return  dto.FindRoomDTO{
			Number: room.Number,
			Type: room.Type,
			PricePerNight: room.PricePerNight,
			Reservations: mappedReservations,
		}
	}).([]dto.FindRoomDTO)

		
	return mapped, nil  
}

