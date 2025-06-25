package service

import (
	"hotel/internal/dto"
	"hotel/internal/logger"
	"hotel/internal/model"
	"hotel/internal/repository"
	"hotel/internal/validator"
	"math"

	"go.uber.org/zap"
)

type RoomService struct {
	roomRepository *repository.RoomRepository
}

func NewRoomService(roomRepository *repository.RoomRepository) *RoomService {
	return  &RoomService{
		roomRepository: roomRepository,
	}
}

func (r * RoomService) CreateRoom(input dto.CreateRoomDTO) (status int, message string) {
	if err := validator.Validate.Struct(input); err != nil {
		logger.ZapLogger.Error(
			"validation error in create room using CreateRoomDTO",
			zap.Error(err),
			zap.String("function", "CreateRoom"),
		)
		return 400, err.Error()	
	}
	input.PricePerNight = math.Round(input.PricePerNight * 100) / 100

	if err := r.roomRepository.CreateRoom(input); err != nil {
		logger.ZapLogger.Error(
			"internal error CreateRoom",
			zap.Error(err),
			zap.String("function", "CreateRoom"),
		)
		return 500, err.Error()
	}

	return 201, "room was created!"
	
}

func (r *RoomService) FindAllRooms() (status int, message interface{}) {
	var rooms []model.Room
	var err error
	if rooms,err = r.roomRepository.FindAllRooms(); err != nil {
		logger.ZapLogger.Error(
			"internal error FindAllRooms",
			zap.Error(err),
			zap.String("function", "FindAllRooms"),
		)
		return 500, err.Error()
	}
	return 200, rooms
}