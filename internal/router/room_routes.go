package router

import (
	"hotel/internal/handler"
	"hotel/internal/logger"
	"hotel/internal/model"
	"hotel/internal/repository"
	"hotel/internal/service"

	"github.com/gofiber/fiber/v2"
)

func setupRoomRoutes(roomRoutes fiber.Router) {
	roomRepo := repository.NewRoomRepository(repository.DB.Model(&model.Room{}))
	roomService := service.NewRoomService(roomRepo)
	roomController := handler.NewRoomController(roomService)
	roomRoutes.Post("/create", roomController.CreateRoom())
	roomRoutes.Get("/all", roomController.FindAllRooms())
	logger.ZapLogger.Info("room's routes are working!")
	
}