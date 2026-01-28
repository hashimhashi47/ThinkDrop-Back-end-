package otprouter

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	"github.com/gofiber/fiber/v2"
)

func OTPRouter(app *fiber.App, OTPController *delivery.AuthControllers) {
	app.Post("/auth/send-otp", OTPController.SentOtp)
	app.Post("/auth/verify-otp", OTPController.VerfiyOtp)
}
