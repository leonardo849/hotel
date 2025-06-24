package router

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/hello", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "welcome!"})
	})
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