package otprouter

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery/otpcontrollers"

	"github.com/gofiber/fiber/v2"
)

func OTPRouter(app *fiber.App, OTPController *otpcontrollers.OtpControllers) {
	app.Post("/auth/send-otp", OTPController.SentOtp)
	app.Post("/auth/verify-otp", OTPController.VerfiyOtp)
}
