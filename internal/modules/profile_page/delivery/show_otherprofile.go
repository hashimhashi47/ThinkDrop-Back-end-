package delivery

import (
	"net/http"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// -> show the others users profile
func (s *ProfileController) ShowOtherUserProfile(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	profileID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	data, err := s.Service.ShowOtherUserProfileService(profileID, UserID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})

}
