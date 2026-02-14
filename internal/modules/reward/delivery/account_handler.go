package delivery

import (
	"net/http"
	"strconv"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

// -> add user bank account
func (s *RewardController) AddUserAcoount(c *fiber.Ctx) error {
	var UserAccountInputs domain.BankAccountInput

	userID, _ := c.Locals("user_id").(uint)

	if err := c.BodyParser(&UserAccountInputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(&UserAccountInputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADGATEWAY, err),
		})
	}

	Data, err := s.service.AddUserAcoountService(userID, UserAccountInputs)
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

// -> withdraw a points to cash
func (s *RewardController) WithdrawPoints(c *fiber.Ctx) error {

	userID, _ := c.Locals("user_id").(uint)

	var WithdrawRequest domain.WithdrawPointsRequest

	if err := c.BodyParser(&WithdrawRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(&WithdrawRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})

	}

	Data, err := s.service.WithdrawPointsService(userID, WithdrawRequest.Points)

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

// -> get the withdraws
func (s *RewardController) GetWithdrawals(c *fiber.Ctx) error {
	userID, _ := c.Locals("user_id").(uint)

	limit, err := strconv.Atoi(c.Query("limit", "5"))

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

	Data, err := s.service.GetWithdrawalsService(userID, limit, offset)

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

// -> get the withdraws with updated data
func (s *RewardController) GetWithdrawalsWithRefersh(c *fiber.Ctx) error {

	userID, _ := c.Locals("user_id").(uint)

	limit, err := strconv.Atoi(c.Query("limit", "5"))

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

	Data, err := s.service.GetWithdrawalsRefershService(userID, limit, offset)

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
