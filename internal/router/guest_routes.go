package router

import (
	"hotel/internal/handler"
	"hotel/internal/logger"
	"hotel/internal/model"
	"hotel/internal/repository"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

func setupGuestRoutes(guestRoutes fiber.Router) {
	modelGuest := repository.DB.Model(&model.Guest{})
	guestRepository  := repository.NewGuestRepository(modelGuest)
	guestService := service.NewGuestService(guestRepository)
	guestController := handler.NewGuestController(guestService)
	guestRoutes.Post("/create", guestController.CreateGuest())
	guestRoutes.Get("/all", guestController.FindAllGuests())
	guestRoutes.Put("/update/:id", guestController.UpdateGuest())
	logger.ZapLogger.Info("guest's routes are running!")
}