package delivery

import (
	"net/http"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

// -> gell all avatars
func (s *ProfileController) GetAvatars(c *fiber.Ctx) error {

	Data, err := s.Service.GetAvatarsService()

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> EditProfile
func (s *ProfileController) EditProfile(c *fiber.Ctx) error {
	var EditProfile domain.EditProfile

	if err := c.BodyParser(&EditProfile); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(&EditProfile); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.EditProfileService(UserID,EditProfile)

if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponseMsg(data, "profile Updated Succesfully"),
	})
}
