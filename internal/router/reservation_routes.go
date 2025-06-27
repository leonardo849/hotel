package router

import (
	"hotel/internal/handler"
	"hotel/internal/repository"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

func setupReservationRoutes(reservationRoutes fiber.Router) {
	reservationRepo := &repository.ReservationRepository{
		GuestRepository: repository.NewGuestRepository(),
		RoomRepository: repository.NewRoomRepository(),
	}
	reservationService := service.NewReservationService(reservationRepo)
	reservationController := handler.NewReservationController(reservationService)
	reservationRoutes.Post("/create", reservationController.CreateReservation())

}