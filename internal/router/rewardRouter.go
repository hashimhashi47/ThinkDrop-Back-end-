package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	rewardmiddileware "thinkdrop-backend/internal/middleware/rewardMiddileware"
	RewardController "thinkdrop-backend/internal/modules/reward/delivery"
	"thinkdrop-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RewardRoutes(app *fiber.App, Redis *redis.Client, db *gorm.DB,
	RewardController *RewardController.RewardController) {

	app.Get("/rewardgetstatus", authmiddileware.AuthenticateMiddileware(Redis, constants.User),
		RewardController.GetRewardStatus)
	app.Post("/createwallet", authmiddileware.AuthenticateMiddileware(Redis, constants.User),
		RewardController.CreateWallet)

	Reward := app.Group("/reward", authmiddileware.AuthenticateMiddileware(Redis, constants.User),
		rewardmiddileware.CheckRewardStatusMiddilware(db))

	Reward.Get("/getwalletdetails", RewardController.GetRewardDetails)
	Reward.Post("/add-bank-account", RewardController.AddUserAcoount)
	Reward.Post("/Withdraw-points", RewardController.WithdrawPoints)
	Reward.Get("/get-withdraws", RewardController.GetWithdrawals)
	Reward.Get("/refresh-withdraws", RewardController.GetWithdrawalsWithRefersh)
}
