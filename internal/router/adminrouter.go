package router

import (
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/admin/delivery"
	"thinkdrop-backend/pkg/constants"

	// "thinkdrop-backend/pkg/constants"

	// "thinkdrop-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/redis/go-redis/v9"
)

func AdminRoutes(app *fiber.App, rds *redis.Client, adminModule *delivery.AdminController) {

	// Create admin group with middleware
	admin := app.Group("/admin", authmiddileware.AuthenticateMiddileware(rds, constants.Admin))

	// HTTP route
	admin.Get("/getstats", adminModule.GetdashboardDetails)
	admin.Get("/getpoststats", adminModule.GetAllusersPost)
	admin.Get("/getaccountstats", adminModule.AddAccountStatus)
	admin.Get("/getaccounts", adminModule.GetWithdrawals)
	admin.Get("/getuserdetails", adminModule.GetUsersDetail)

	admin.Post("/block-user/:id", adminModule.BlockUser)
	admin.Post("/unblock-user/:id", adminModule.UnBlockUser)

	admin.Get("/getflagedpost", adminModule.GetAllFlagedPost)
	admin.Get("/getwallets", adminModule.GetWallets)
	admin.Post("/block-wallet/:id", adminModule.BlockWallet)
	admin.Post("/unblock-wallet/:id", adminModule.UnBlockWallet)
	admin.Post("/verify-bank-account/:id", adminModule.VerifyAccount)
	admin.Delete("/delete-posts/:id", adminModule.RemoveTheFlaggedPost)
	admin.Post("/safe-posts/:id", adminModule.SafePosting)
	admin.Get("/profile", adminModule.GetProfile)
	admin.Put("/updateprofile", adminModule.UpadteProfile)
	admin.Put("/updateuser/:id", adminModule.UpdateUserProfile)
	admin.Post("/adduser", adminModule.AddUser)
	admin.Delete("/deleteuser/:id", adminModule.DeleteUser)

	admin.Post("/createinterest", adminModule.AddMainIntrest)
	admin.Post("/createsubinterest", adminModule.AddSubIntrest)
	admin.Put("/updateinterest/:id", adminModule.UpadteMainIntrest)
	admin.Put("/updatesubinterest/:id", adminModule.UpdateSubIntrest)
	admin.Delete("/deleteinterest/:id", adminModule.DeleteIntrest)
	admin.Delete("/deletesubinterest/:id", adminModule.DeleteSubIntrest)

	admin.Post("/block-post/:id", adminModule.BlockPost)
	admin.Post("/unblock-post/:id", adminModule.UnBlockPost)
	admin.Put("/updatepostinterest/:id", adminModule.UpdatePostIntrests)

	admin.Get("/complaints", adminModule.GetAllComplaints)
	admin.Put("/complaints/:id", adminModule.ConsiderTheIssue)

	// WebSocket route
	app.Get("/admin/ws", websocket.New(adminModule.Handle))
}
