package router

import (
	"log"
	"os"
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/chat/delivery"
	Token "thinkdrop-backend/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func ChatRoutes(app *fiber.App, rds *redis.Client, chatModule *delivery.ChatHandler) {

	//web socket route for connecting
	app.Get("/ws/chat", websocket.New(func(c *websocket.Conn) {

		token := c.Query("token")
		log.Println("🕹️Log", token)
		if token == "" {
			token = c.Cookies("Access_token")
			log.Println("🕹️Log", token)
		}
		log.Println("🕹️Log", token)
		if token == "" {
			log.Println("token missing")
			c.Close()
			return
		}

		KEY := []byte(os.Getenv("JWT_SECRET_KEY"))

		claim := &Token.Claims{}
		parsedToken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
			return KEY, nil
		})

		if err != nil || !parsedToken.Valid {
			log.Println("invalid token:", err)
			c.Close()
			return
		}

		chatModule.HandleWS(c, claim.UserId)
	}))

	//http routes
	app.Get("/ws/getsidebar", authmiddileware.AuthenticateMiddileware(rds), chatModule.GetChats)

	//get ful chat on a particular person by id
	app.Get("/ws/chat/:id/messages",
		authmiddileware.AuthenticateMiddileware(rds),
		chatModule.GetChatMessages)

		
	//to start a particular chat
	app.Post("/ws/conversation",
		authmiddileware.AuthenticateMiddileware(rds),
		chatModule.StartConversation)

}
