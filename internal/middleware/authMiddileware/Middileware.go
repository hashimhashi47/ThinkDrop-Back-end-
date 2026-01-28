package authmiddileware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"os"
	token "thinkdrop-backend/pkg/jwt"
)

func AuthenticateMiddileware(rds *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		AccessToken := c.Cookies("Access_token")
		KEY := []byte(os.Getenv("JWT_SECRET_KEY"))

		if AccessToken == "" {
			return fiber.ErrUnauthorized
		}

		claim := &token.Claims{}
		tkn, err := jwt.ParseWithClaims(AccessToken, claim, func(t *jwt.Token) (interface{}, error) {
			return KEY, nil
		})

		if err != nil {
			return fiber.ErrUnauthorized
		}

		if !tkn.Valid {
			return fiber.ErrUnauthorized
		}
		c.Locals("user_id", claim.UserId)
		return c.Next()
	}
}
