package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/profile_page/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ProfileRoute(app *fiber.App, rds *redis.Client, ProfileController *delivery.ProfileController) {
	User := app.Group("/users", authmiddileware.AuthenticateMiddileware(rds))

	User.Get("/avatars", ProfileController.GetAvatars)
	User.Put("/updateprofile", ProfileController.EditProfile)
	User.Get("/profile", ProfileController.ShowProfile)
	User.Get("/writings", ProfileController.GetAllWritings)
	User.Get("/followings", ProfileController.GetFollowings)
	User.Get("/followers", ProfileController.GetAllFollowers)

	User.Get("/intrst", ProfileController.GetUserIntrest)
	User.Get("/:id", ProfileController.ShowOtherUserProfile)
	User.Post("/follow/:id", ProfileController.FollowUser)
	User.Post("/unfollow/:id", ProfileController.UserUnfollow)
	
}
