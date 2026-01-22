package router

import (
	"github.com/gofiber/fiber/v2"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
)

// → Route registrations
func UserRoutes(app *fiber.App, controllers *delivery.UserController) {
	app.Post("/signup", controllers.UserSignup)
}
