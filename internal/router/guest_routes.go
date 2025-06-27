package router

import (
	"hotel/internal/handler"
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

func setupGuestRoutes(guestRoutes fiber.Router) {
	guestRepository  := repository.NewGuestRepository()
	guestService := service.NewGuestService(guestRepository)
	guestController := handler.NewGuestController(guestService)
	guestRoutes.Post("/create", guestController.CreateGuest())
	guestRoutes.Get("/all", guestController.FindAllGuests())
	guestRoutes.Put("/update/:id", guestController.UpdateGuest())
	guestRoutes.Delete("/delete/:id", guestController.DeleteGuest())
	logger.ZapLogger.Info("guest's routes are running!")
}