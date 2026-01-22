package delivery

import (
	"net/http"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/entity"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	"thinkdrop-backend/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

//→ Controllers (HTTP handlers)

type UserController struct {
	services *usecase.UserServices
}

func NewUserController(s *usecase.UserServices) *UserController {
	return &UserController{services: s}
}

// ->User Signup binding and sent to controllers
func (s *UserController) UserSignup(c *fiber.Ctx) error {
	var uservalidate entity.UserValidate

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

	data, err := s.services.UserLoginService(&uservalidate)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})

}
