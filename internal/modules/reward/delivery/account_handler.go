package delivery

import (
	"net/http"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

// -> add user baank account

func (s *RewardController) AddUserAcoount(c *fiber.Ctx) error {
	var UserAccountInputs domain.BankAccountInput

	userID, _ := c.Locals("user_id").(uint)

	if err := c.BodyParser(&UserAccountInputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADGATEWAY, err),
		})
	}

	if err := validator.Validate.Struct(&UserAccountInputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADGATEWAY, err),
		})
	}

	Data, err := s.service.AddUserAcoountService(userID, UserAccountInputs)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.INTERNALSERVERERROR, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})

}
