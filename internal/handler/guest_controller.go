package handler

import (
	"hotel/internal/dto"
	"hotel/internal/helper"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

type GuestController struct {
	guestService *service.GuestService
}

func NewGuestController(guestService *service.GuestService) *GuestController {
	return  &GuestController{
		guestService: guestService,
	}
}

// CreateGuest godoc
// @Description Create a new guest
// @Tags guests
// @Accept json
// @Produce json
// @Param guest body dto.CreateGuestDTO true "Guest data"
// @Success 201 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Router /guest/create [post]
func (g *GuestController) CreateGuest() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input dto.CreateGuestDTO
		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := g.guestService.CreateGuest(input)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property: message})
	}
}


// FindGuests godoc
// @Description Find all guests
// @Tags guests
// @Accept json
// @Produce json
// @Success 200 {array} dto.FindGuestDTO
// @Failure 400 {object} errorResponse
// @Router /guest/all [get]
func (g *GuestController) FindAllGuests() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		status, message := g.guestService.FindAllGuests()
		if status >= 400 {
			return  ctx.Status(status).JSON(fiber.Map{"error": message})
		}
		return  ctx.Status(status).JSON(message)
	}
}


// UpdatGuests godoc
// @Description update guest
// @Tags guests
// @Accept json
// @Produce json
// @Success 200 {array} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /guest/update/{id} [put]
func (g *GuestController) UpdateGuest() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var input dto.UpdateGuestDTO
		if err := ctx.BodyParser(&input); err != nil {
			return  ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		status, message := g.guestService.UpdateGuest(id, input)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property:message})
	}
}

// DeleteGuests godoc
// @Description delete guest
// @Tags guests
// @Accept json
// @Produce json
// @Success 200 {array} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /guest/delete/{id} [delete]
func (g *GuestController) DeleteGuest() fiber.Handler {
	return  func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		status, message := g.guestService.DeleteGuest(id)
		var property string
		helper.SetProperty(&property, status)
		return  ctx.Status(status).JSON(fiber.Map{property:message})
	}
}