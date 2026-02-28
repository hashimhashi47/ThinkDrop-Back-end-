package delivery

import (
	"net/http"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"
	validator "thinkdrop-backend/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

func (a *ProfileController) RaiseComplaint(c *fiber.Ctx) error {
	var Complaint domain.CreateReportRequest
	UserID, _ := c.Locals("user_id").(uint)

	if err := c.BodyParser(&Complaint); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := validator.Validate.Struct(Complaint); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADGATEWAY, err),
		})
	}

	if err := a.Service.RaiseComplaintService(Complaint,UserID); err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("Successfully sent the report"),
	})
}
