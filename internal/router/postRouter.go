package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/post/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func PostRoutes(app *fiber.App, PostController *delivery.PostControllers, rds *redis.Client) {
	Post := app.Group("/post", authmiddileware.AuthenticateMiddileware(rds))

	Post.Post("/uploadpost", PostController.AddPost)
	Post.Get("/getalluserwithposts", PostController.ShowPosts)
	Post.Get("/getallfeed", PostController.Userfeed)
}
