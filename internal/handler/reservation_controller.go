package handler

import (
	"hotel/internal/dto"
	"hotel/internal/helper"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ReservationController struct {
	reservationService *service.ReservationService
}

func NewReservationController(reservationService *service.ReservationService) *ReservationController {
	return  &ReservationController{
		reservationService: reservationService,
	}
}

func (r *ReservationController) CreateReservation() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input dto.CreateReservationDTO
		if err := ctx.BodyParser(&input); err != nil {
			return  ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := r.reservationService.CreateReservation(input)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property:message})
	}
}
