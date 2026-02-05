package delivery

import (
	"net/http"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// -> show the others users profile
func (s *ProfileController) ShowOtherUserProfile(c *fiber.Ctx) error {
	profileID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	data, err := s.Service.ShowOtherUserProfileService(profileID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})

}