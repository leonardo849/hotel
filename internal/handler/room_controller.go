package handler

import (
	"hotel/internal/dto"
	"hotel/internal/helper"
	"hotel/internal/logger"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
// @Router /room/create [post]
func (r *RoomController) CreateRoom() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input dto.CreateRoomDTO
		if err := ctx.BodyParser(&input); err != nil {
			logger.ZapLogger.Error(
				"error body parser create room",
				zap.Error(err),
				zap.String("function", "create room controller"),
			)
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
// @Success 200 {array} []dto.FindRoomDTO
// @Failure 400 {object} errorResponse
// @Router /room/all [get]
func (r *RoomController) FindAllRooms() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		status, message := r.roomService.FindAllRooms()
		if status >= 400 {
			return  ctx.Status(status).JSON(fiber.Map{"error": message})
		}
		return  ctx.Status(status).JSON(message)
	}
}



// FindOneRoom godoc
// @Description find one room by id
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "room id"
// @Success 200 {object} dto.FindRoomDTO
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /room/{id} [get]
func (r *RoomController) FindOneRoom() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		status, message := r.roomService.FindOneRoom(id)
		if status >= 400 {
			return  ctx.Status(status).JSON(fiber.Map{"error": message})
		}
		return  ctx.Status(200).JSON(message)
	}
}


// UpdateRoom godoc
// @Description update room by id
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "room id"
// @Success 200 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /room/update/{id} [put]
func (r *RoomController) UpdateRoom() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var input dto.UpdateRoomDTO
		if err := ctx.BodyParser(&input); err != nil {
			logger.ZapLogger.Error(
				"error body parser update room",
				zap.Error(err),
				zap.String("function", "update room controller"),
			)
			return  ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := r.roomService.UpdateRoom(id, input)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property:message})

	}
}


// DeletRoom godoc
// @Description delete room by id
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "room id"
// @Success 200 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /room/delete/{id} [delete]
func (r *RoomController) DeleteRoom() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		status, message := r.roomService.DeleteRoom(id)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property:message})
	}
}