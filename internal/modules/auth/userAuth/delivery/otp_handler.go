package delivery

import (
	"errors"
	"net/http"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

// ->sent OTP Controller
func (s *AuthControllers) SentOtp(c *fiber.Ctx) error {
	var OtpEmail struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := c.BodyParser(&OtpEmail); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(OtpEmail); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	OTP, err := s.services.SentOtpService(OtpEmail.Email)

	if errors.Is(err, errors.New("Request limit exceeded, wait for 10 min")) {
		return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.TOOMANYREQUESTS, err),
		})
	}

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponseMsg(OTP, "OTP valid upto 5 min"),
	})
}

// -> Verify the OTP is valid or Not

func (s *AuthControllers) VerfiyOtp(c *fiber.Ctx) error {
	var VerifyOtp struct {
		Email string `json:"email" validate:"required,email"`
		Otp   string `json:"otp" validate:"required,len=6,numeric"`
	}

	if err := c.BodyParser(&VerifyOtp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(&VerifyOtp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := s.services.OTPverifyService(VerifyOtp.Email, VerifyOtp.Otp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("OTP valid upto 5 min"),
	})
}
