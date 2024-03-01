package main

import (
	"track_flight_api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New())

	app.Get("/calculate", handler.TrackHandler)

	app.Listen(":8080")
}
