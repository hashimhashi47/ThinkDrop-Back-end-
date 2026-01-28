package delivery

import (
	"net/http"
	InterestService "thinkdrop-backend/internal/modules/interest/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type InterestControllers struct {
	service *InterestService.InterestService
}

func NewInterestControllers(s *InterestService.InterestService) *InterestControllers {
	return &InterestControllers{service: s}
}

func (s *InterestControllers) ShowIntrests(c *fiber.Ctx) error {
	interest, err := s.service.ShowIntrestsService()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"interests": response.SuccessResponse(interest),
	})
}
