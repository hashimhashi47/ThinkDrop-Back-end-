package delivery

import (
	"net/http"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"
	"time"

	"github.com/gofiber/fiber/v2"
)

//→ Controllers (HTTP handlers)

type AuthControllers struct {
	services *usecase.AuthService
}

func NewUserController(s *usecase.AuthService) *AuthControllers {
	return &AuthControllers{services: s}
}

// -> User Signup with bindings and sent to services
func (s *AuthControllers) UserSignup(c *fiber.Ctx) error {
	var uservalidate domain.UserValidate

	if err := c.BodyParser(&uservalidate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(uservalidate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADGATEWAY, err),
		})
	}

	data, err := s.services.UserSignupService(&uservalidate)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

// -> User Login with bindings and sent to services
func (s *AuthControllers) UserLogin(c *fiber.Ctx) error {
	var validateLogin domain.Login

	if err := c.BodyParser(&validateLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(validateLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, AccessToken, RefershToken, err := s.services.UserLoginService(&validateLogin)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Access_token",
		Value:    AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "Refersh_token",
		Value:    RefershToken,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	Result := map[string]string{
		"Email":       Data.Email,
		"Name":        Data.FullName,
		"AccessToken": AccessToken,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Result),
	})
}
