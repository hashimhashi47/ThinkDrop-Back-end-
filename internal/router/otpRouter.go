package router

import (
	"github.com/gofiber/fiber/v2"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
)

func OTPRouter(app *fiber.App, OTPController *delivery.AuthControllers) {
	Auth := app.Group("/auth")
	Auth.Post("/send-otp", OTPController.SentOtp)
	Auth.Post("/verify-otp", OTPController.VerfiyOtp)
}
