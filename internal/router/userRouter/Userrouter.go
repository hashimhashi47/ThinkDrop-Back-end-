package userrouter

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	"github.com/gofiber/fiber/v2"
)

// → Route registrations
func UserRoutes(app *fiber.App, controllers *delivery.AuthControllers) {
	app.Post("/signup", controllers.UserSignup)
	app.Post("/login", controllers.UserLogin)
}
