package delivery

import (
	"errors"
	"net/http"
	"thinkdrop-backend/pkg/constants"
	"time"

	"github.com/gofiber/fiber/v2"
)

// -> Refresh token handler it will handle the 404 situat
func (s *AuthControllers) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("Refersh_token")

	token, err := s.services.RefereshTokenService(refreshToken)

	if err != nil {
		if errors.Is(err, constants.ErrTokenExpired) {
			return fiber.ErrUnauthorized
		}
		if errors.Is(err, constants.ErrTokenMismatch) {
			return fiber.ErrUnauthorized
		}
		if errors.Is(err, constants.ErrTokenNotFound) {
			return fiber.ErrUnauthorized
		}
		return fiber.ErrInternalServerError
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Access_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
	})

	return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}
