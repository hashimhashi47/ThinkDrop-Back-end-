package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/post/delivery"
	"thinkdrop-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func PostRoutes(app *fiber.App, PostController *delivery.PostControllers, rds *redis.Client) {
	Post := app.Group("/post", authmiddileware.AuthenticateMiddileware(rds, constants.User))

	Post.Post("/uploadpost", PostController.AddPost)
	Post.Get("/getalluserwithposts", PostController.ShowPosts)
	Post.Get("/getallfeed", PostController.Userfeed)
	Post.Post("/likepost/:id", PostController.LikePost)
	Post.Post("/unlikepost/:id", PostController.UnLikePost)
	Post.Post("/reportpost", PostController.ReportPost)
}
