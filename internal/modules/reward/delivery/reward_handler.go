package delivery

import (
	"fmt"
	"net/http"
	RewardService "thinkdrop-backend/internal/modules/reward/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type RewardController struct {
	service *RewardService.RewardService
}

func NewRewardController(s *RewardService.RewardService) *RewardController {
	return &RewardController{service: s}
}

// -> get the wallet status
func (s *RewardController) GetRewardStatus(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)

	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	hasWallet, err := s.service.GetRewardStatusService(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.INTERNALSERVERERROR, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponseMsg(hasWallet, "hasWallet"),
	})
}

// -> wallet creating for the specific user
func (s *RewardController) CreateWallet(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	Data, err := s.service.CreateWalletService(UserID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.INTERNALSERVERERROR, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponseMsg(Data, "Wallet Created succesfully"),
	})
}

// -> get the reward details
func (s *RewardController) GetRewardDetails(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)
	fmt.Println(UserID)
	Data, err := s.service.GetRewardDetailsService(UserID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.INTERNALSERVERERROR, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}
