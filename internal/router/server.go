package router

import (
	"fmt"
	"hotel/internal/logger"
	"hotel/internal/middleware"
	"os"
	_ "hotel/docs" 
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Hotel API
func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/hello", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "welcome!"})
	})
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(middleware.LogRequestsMiddleware())
	roomGroup := app.Group("/room")
	guestGroup := app.Group("/guest")
	reservationGroup := app.Group("/reservation")
	setupReservationRoutes(reservationGroup)
	setupRoomRoutes(roomGroup)
	setupGuestRoutes(guestGroup)
	logger.ZapLogger.Info("app is ready!")
	return app
}

func RunServer() error {
	app := SetupApp()

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "welcome!"})
	})

	return app.Listen(":" + port)
}