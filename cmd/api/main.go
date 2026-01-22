package main

import (
	"thinkdrop-backend/internal/router"

	"github.com/gofiber/fiber/v2"
)

// → App entry point (starts server, loads config)
func main() {
	app := fiber.New()
	router.UserRoutes(app)
	app.Listen(":8000")
}
