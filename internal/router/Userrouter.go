package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	AuthDelivery "thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	InterestDelivery "thinkdrop-backend/internal/modules/interest/delivery"
	"thinkdrop-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// → Route registrations
func UserRoutes(app *fiber.App, Redis *redis.Client, AuthControllers *AuthDelivery.AuthControllers,
	IntrestControllers *InterestDelivery.InterestControllers) {
	Auth := app.Group("/auth")
	Auth.Post("/signup", AuthControllers.UserSignup)
	Auth.Post("/login", AuthControllers.UserLogin)
	Auth.Post("/refersh", AuthControllers.RefreshToken)
	Auth.Post("/logout", authmiddileware.AuthenticateMiddileware(Redis,constants.User), AuthControllers.Logout)

	app.Get("/getallinterest", IntrestControllers.ShowIntrests)
	app.Post("/addinterest", authmiddileware.AuthenticateMiddileware(Redis,constants.User), IntrestControllers.UserAddIntersts)
}
