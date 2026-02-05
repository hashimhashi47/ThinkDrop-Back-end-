package rewardmiddileware

import (
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CheckRewardStatusMiddilware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		UserID, _ := c.Locals("user_id").(uint)

		var Wallet domain.Wallet

		if err := db.Where("user_id = ?", UserID).First(&Wallet).Error; err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "wallet not found",
			})
		}

		if Wallet.IsWalletActive == constants.WalletBlocked {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "wallet is blocked",
			})
		}

		c.Locals("wallet", Wallet)
		return c.Next()
	}
}
