package userrouter

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	AuthDelivery "thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	InterestDelivery "thinkdrop-backend/internal/modules/interest/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// → Route registrations
func UserRoutes(app *fiber.App, Redis *redis.Client, AuthControllers *AuthDelivery.AuthControllers,
	IntrestControllers *InterestDelivery.InterestControllers) {

	app.Post("/signup", AuthControllers.UserSignup)
	app.Post("/login", AuthControllers.UserLogin)
	app.Post("/auth/refersh", AuthControllers.RefreshToken)
	app.Get("/getallinterest", IntrestControllers.ShowIntrests)
	app.Post("/addinterest", authmiddileware.AuthenticateMiddileware(Redis), IntrestControllers.UserAddIntersts)
}
