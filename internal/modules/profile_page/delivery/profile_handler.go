package delivery

import (
	"net/http"
	"strconv"
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

// -> show own profile of the user
func (s *ProfileController) ShowProfile(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.ShowProfileService(UserID)

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

// -> follow a user
func (s *ProfileController) FollowUser(c *fiber.Ctx) error {
	profileID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	UserID, _ := c.Locals("user_id").(uint)

	data1, data2, err := s.Service.FollowUserService(UserID, profileID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
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
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
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

	limit, err := strconv.Atoi(c.Query("limit", "10"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: "invalid limit",
		})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: "invalid offset",
		})
	}

	data, err := s.Service.GetAllWritingsService(UserID, limit, offset)

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

// -> get all followers of the user
func (s *ProfileController) GetAllFollowers(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.GetAllFollowersService(UserID)

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

// -> get all followings of the user
func (s *ProfileController) GetFollowings(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	data, err := s.Service.GetAllFollowingService(UserID)

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

// -> get user intrests
func (s *ProfileController) GetUserIntrest(c *fiber.Ctx) error {
	userIDAny := c.Locals("user_id")
	userID, ok := userIDAny.(uint)
	if !ok || userID == 0 {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or missing user id",
		})
	}

	data, err := s.Service.GetUserIntrest(userID)
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
