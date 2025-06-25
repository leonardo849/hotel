package handler

import (
	"hotel/internal/dto"
	"hotel/internal/helper"
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
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property: message})
	}
}

func (r *RoomController) FindAllRooms() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		status, message := r.roomService.FindAllRooms()
		if status >= 400 {
			return  ctx.Status(status).JSON(fiber.Map{"error": message})
		}
		return  ctx.Status(status).JSON(message)
	}
}