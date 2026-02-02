package delivery

import (
	"net/http"
	Profileservice "thinkdrop-backend/internal/modules/profile_page/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	Service *Profileservice.ProfileService
}

func NewProfileControllers(s *Profileservice.ProfileService) *ProfileController {
	return &ProfileController{Service: s}
}

func (s *ProfileController) ShowProfile(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.ShowProfileService(UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

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
