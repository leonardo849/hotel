package repository

import (
	"hotel/internal/dto"
	"hotel/internal/model"
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

func (r * RoomRepository) CreateRoom(input dto.CreateRoomDTO) error {
	room := model.Room{
		Number: input.Number,
		Type: input.Type,
		PricePerNight: input.PricePerNight,
	}
	return r.roomModel.Create(&room).Error
}

