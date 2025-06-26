package service

import (
	"hotel/internal/dto"
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/validator"
	"math"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RoomService struct {
	roomRepository *repository.RoomRepository
}

func NewRoomService(roomRepository *repository.RoomRepository) *RoomService {
	return &RoomService{
		roomRepository: roomRepository,
	}
}

func (r *RoomService) CreateRoom(input dto.CreateRoomDTO) (status int, message string) {
	if err := validator.Validate.Struct(input); err != nil {
		logger.ZapLogger.Error(
			"validation error in create room using CreateRoomDTO",
			zap.Error(err),
			zap.String("function", "CreateRoom"),
		)
		return 400, err.Error()
	}
	input.PricePerNight = math.Round(input.PricePerNight*100) / 100

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
	var rooms []dto.FindRoomDTO
	var err error
	if rooms, err = r.roomRepository.FindAllRooms(); err != nil {
		logger.ZapLogger.Error(
			"internal error FindAllRooms",
			zap.Error(err),
			zap.String("function", "FindAllRooms"),
		)
		return 500, err.Error()
	}
	return 200, rooms
}

func (r *RoomService) FindOneRoom(id string) (status int, message interface{}) {
	var room *dto.FindRoomDTO
	_, err := uuid.Parse(id)
	if err != nil {
		logger.ZapLogger.Error(
			"bad request uuid is invalid",
			zap.Error(err),
			zap.String("function", "findOneRoom"),
		)
		return 400, "uuid is invalid"
	}
	if room, err = r.roomRepository.FindOneRoom(id); err != nil {
		logger.ZapLogger.Error(
			"room was not found",
			zap.Error(err),
			zap.String("function", "findOneRoom"),
		)
		return 404, "room wasn't found"
	}
	return 200, &room
}

func (r *RoomService) UpdateRoom(id string, input dto.UpdateRoomDTO) (status int, message interface{}) {
	_, err := uuid.Parse(id)
	if err != nil {
		logger.ZapLogger.Error(
			"bad request uuid is invalid",
			zap.Error(err),
			zap.String("function", "findOneRoom"),
		)
		return 400, "uuid is invalid"
	}
	status, res := r.FindOneRoom(id)
	if status >= 400 {
		return status, res
	}
	return 200, "room was updated"
}

func (r *RoomService) DeleteRoom(id string) (status int, message interface{}) {
	_, err := uuid.Parse(id)
	if err != nil {
		logger.ZapLogger.Error(
			"bad request uuid is invalid",
			zap.Error(err),
			zap.String("function", "findOneRoom"),
		)
		return 400, "uuid is invalid"
	}
	status, res := r.FindOneRoom(id)
	if status >= 400 {
		return status, res
	}
	return 200, "room was deleted"
}
