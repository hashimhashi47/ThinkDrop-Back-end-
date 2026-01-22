package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"sucess": "ok"})
}

// → Route registrations
func UserRoutes(app *fiber.App) {
	
}
