package userrouter

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery/authcontrollers"

	"github.com/gofiber/fiber/v2"
)

// → Route registrations
func UserRoutes(app *fiber.App, controllers *authcontrollers.UserController) {
	app.Post("/signup", controllers.UserSignup)
	app.Post("/login", controllers.UserLogin)
}
