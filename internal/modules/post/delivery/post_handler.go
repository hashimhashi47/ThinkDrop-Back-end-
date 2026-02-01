package delivery

import (
	"net/http"
	"strconv"
	domain "thinkdrop-backend/internal/Common"
	PostService "thinkdrop-backend/internal/modules/post/usecase"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type PostControllers struct {
	Service *PostService.PostService
}

func NewPostControllers(s *PostService.PostService) *PostControllers {
	return &PostControllers{Service: s}
}

// -> Add post controller will bind the deatils and pass to service layer
func (s *PostControllers) AddPost(c *fiber.Ctx) error {
	var PostDetails domain.Post

	if err := c.BodyParser(&PostDetails); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	UserID, _ := c.Locals("user_id").(uint)
	data, err := s.Service.AddPostService(PostDetails, UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	resp := map[string]interface{}{
		"content":      data.Content,
		"created_date": data.CreatedAt.Format("2006-01-02"),
		"created_time": data.CreatedAt.Format("15:04:05"),
		"likecount":    data.LikeCount,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(resp),
	})
}

// -> show all the posts on thier profile
func (r *PostControllers) ShowPosts(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)
	data, err := r.Service.ShowPostsServices(UserID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})

}

// -> show the feed with releted to their intrest
func (s *PostControllers) Userfeed(c *fiber.Ctx) error {
	UserID, _ := c.Locals("user_id").(uint)

	limit, err := strconv.Atoi(c.Query("limit", "20"))

	if err != nil {
		return c.Status(constants.BADREQUEST).JSON(fiber.Map{
			constants.Error: "invalid limit",
		})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))

	if err != nil {
		return c.Status(constants.BADREQUEST).JSON(fiber.Map{
			constants.Error: "invalid offset",
		})
	}

	Data, err := s.Service.UserFeedService(UserID, limit, offset)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}
