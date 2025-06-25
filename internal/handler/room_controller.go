package handler

import (
	"hotel/internal/dto"
	"hotel/internal/helper"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

type messageResponse struct {
    Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}
type RoomController struct {
	roomService *service.RoomService
}

func NewRoomController(roomService *service.RoomService) *RoomController {
	return &RoomController{
		roomService: roomService,
	}
}

// CreateRoom godoc
// @Description Create a new hotel room
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body dto.CreateRoomDTO true "Room data"
// @Success 201 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Router /rooms/create [post]
func (r *RoomController) CreateRoom() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input dto.CreateRoomDTO
		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := r.roomService.CreateRoom(input)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property: message})
	}
}


// FindAllRooms godoc
// @Description find all rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {array} model.Room
// @Failure 400 {object} errorResponse
// @Router /rooms/all [get]
func (r *RoomController) FindAllRooms() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		status, message := r.roomService.FindAllRooms()
		if status >= 400 {
			return  ctx.Status(status).JSON(fiber.Map{"error": message})
		}
		return  ctx.Status(status).JSON(message)
	}
}