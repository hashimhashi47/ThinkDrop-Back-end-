package delivery

import (
	"net/http"
	domain "thinkdrop-backend/internal/Common"
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

// -> get enitire interests to show for selecting users
func (s *InterestControllers) ShowIntrests(c *fiber.Ctx) error {
	interest, err := s.service.ShowIntrestsService()

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"interests": response.SuccessResponse(interest),
	})
}

// -> user will select the intrests it will added on database
func (s *InterestControllers) UserAddIntersts(c *fiber.Ctx) error {
	var Req domain.Req

	if err := c.BodyParser(&Req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	ID, _ := c.Locals("user_id").(uint)

	Data, err := s.service.UserAddInterstsService(ID, Req)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	result := map[string]interface{}{
		"username": Data.FullName,
		"Intersts": Data.SelectedSubs,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(result),
	})

}
