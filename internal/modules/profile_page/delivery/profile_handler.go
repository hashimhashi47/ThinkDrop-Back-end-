package delivery

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	Profileservice "thinkdrop-backend/internal/modules/profile_page/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
)

type ProfileController struct {
	Service *Profileservice.ProfileService
}

func NewProfileControllers(s *Profileservice.ProfileService) *ProfileController {
	return &ProfileController{Service: s}
}

// -> show own profile of the user
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

// -> follow a user
func (s *ProfileController) FollowUser(c *fiber.Ctx) error {
	profileID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	UserID, _ := c.Locals("user_id").(uint)

	data1, data2, err := s.Service.FollowUserService(UserID, profileID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	result := map[string]interface{}{
		"User":         data1,
		"FollowedUser": data2,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(result),
	})
}

// -> unfollow a user
func (s *ProfileController) UserUnfollow(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	profileID := uint(id)

	if err != nil {
		return fiber.ErrBadRequest
	}

	UserID, _ := c.Locals("user_id").(uint)

	data1, data2, err := s.Service.UnfollowUser(UserID, profileID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	result := map[string]interface{}{
		"User":         data1,
		"FollowedUser": data2,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(result),
	})

}

// -> gell all wrtitings of the user
func (s *ProfileController) GetAllWritings(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.GetAllWritingsService(UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

// -> get all followers of the user
func (s *ProfileController) GetAllFollowers(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.GetAllFollowersService(UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

// -> get all followings of the user
func (s *ProfileController) GetFollowings(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.GetAllFollowingService(UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}
