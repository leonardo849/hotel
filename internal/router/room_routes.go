package router

import (
	"hotel/internal/handler"
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

func setupRoomRoutes(roomRoutes fiber.Router) {
	roomRepo := repository.NewRoomRepository()
	roomService := service.NewRoomService(roomRepo)
	roomController := handler.NewRoomController(roomService)
	roomRoutes.Post("/create", roomController.CreateRoom())
	roomRoutes.Get("/all", roomController.FindAllRooms())
	roomRoutes.Get("/one/:id", roomController.FindOneRoom())
	roomRoutes.Put("/update/:id", roomController.UpdateRoom())
	roomRoutes.Delete("/delete/:id", roomController.DeleteRoom())
	logger.ZapLogger.Info("room's routes are working!")
	
}