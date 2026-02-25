package authmiddileware

import (
	"log"
	"os"
	token "thinkdrop-backend/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func AuthenticateMiddileware(rds *redis.Client, roles ...string) fiber.Handler {
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
		c.Locals("role", claim.Role)
		log.Println(claim.Role)
		// RBAC(Role based access controll)
		if len(roles) > 0 {
			allowed := false

			for _, v := range roles {
				if claim.Role == v {
					allowed = true
					break
				}
			}
			if !allowed {
				return fiber.ErrForbidden
			}
		}
		return c.Next()

	}
}
