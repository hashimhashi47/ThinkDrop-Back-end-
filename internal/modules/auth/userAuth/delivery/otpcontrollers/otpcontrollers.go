package otpcontrollers

import (
	"net/http"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase/otpService"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

type OtpControllers struct {
	service *otpservice.OtpService
}

func NewOtpServices(s *otpservice.OtpService) *OtpControllers {
	return &OtpControllers{service: s}
}

// ->sent OTP Controller
func (s *OtpControllers) SentOtp(c *fiber.Ctx) error {
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

	OTP, err := s.service.SentOtpService(OtpEmail.Email)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponseMsg(OTP, "OTP valid upto 5 min"),
	})
}
