package userrouter

import (
	AuthDelivery "thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	InterestDelivery "thinkdrop-backend/internal/modules/interest/delivery"

	"github.com/gofiber/fiber/v2"
)

// → Route registrations
func UserRoutes(app *fiber.App, AuthControllers *AuthDelivery.AuthControllers,
	IntrestControllers *InterestDelivery.InterestControllers) {

	app.Post("/signup", AuthControllers.UserSignup)
	app.Post("/login", AuthControllers.UserLogin)
	app.Post("/auth/refersh", AuthControllers.RefreshToken)
	app.Get("/getallinterest", IntrestControllers.ShowIntrests)
}
