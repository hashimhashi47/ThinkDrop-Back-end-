package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/chat/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/redis/go-redis/v9"
)

func ChatRoutes(app *fiber.App, rds *redis.Client, chatModule *delivery.ChatHandler) {
	app.Get("/ws/chat", websocket.New(chatModule.HandleWS))
	app.Get("/ws/getsidebar", authmiddileware.AuthenticateMiddileware(rds), chatModule.GetChats)
}
