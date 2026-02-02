package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/profile_page/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ProfileRoute(app *fiber.App, rds *redis.Client, ProfileController *delivery.ProfileController) {
	app.Get("/users/:id", authmiddileware.AuthenticateMiddileware(rds), ProfileController.ShowOtherUserProfile)
}
