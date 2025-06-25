package handler

import (
	"hotel/internal/dto"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RoomController struct {
	roomService *service.RoomService
}

func NewRoomController(roomService *service.RoomService) *RoomController {
	return &RoomController{
		roomService: roomService,
	}
}

func (r *RoomController) CreateRoom() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input dto.CreateRoomDTO
		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := r.roomService.CreateRoom(input)
		property := "message"
		if status >= 400 {
			property = "error"
		}
		return  ctx.Status(status).JSON(fiber.Map{property: message})
	}
}