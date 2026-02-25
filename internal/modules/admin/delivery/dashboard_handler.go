package delivery

import (
	"net/http"
	"strconv"
	"thinkdrop-backend/internal/modules/admin/domain"
	"thinkdrop-backend/internal/modules/admin/usecase"
	"thinkdrop-backend/internal/modules/admin/websocket"
	"thinkdrop-backend/pkg/constants"
	"thinkdrop-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	service *usecase.AdminService
	hub     *websocket.Hub
}

func NewAdminController(s *usecase.AdminService, h *websocket.Hub) *AdminController {
	return &AdminController{service: s, hub: h}
}

// -> get tha statiscs for admin dashboard handler
func (a *AdminController) GetdashboardDetails(c *fiber.Ctx) error {

	data, err := a.service.GetdashboardDetailsService()

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

// -> get all users posts and details
func (a *AdminController) GetAllusersPost(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	offset, err := strconv.Atoi(c.Query("offset", "0"))

	Data, total, err := a.service.GetAllusersPostService(limit, offset)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: fiber.Map{
			"data":  Data,
			"total": total,
		},
	})
}

// -> get the accounts status with total points the wholw users have and how they are redeemend
func (r *AdminController) AddAccountStatus(c *fiber.Ctx) error {

	Data, err := r.service.AddAccountStatusService()

	if err != nil {
		status := response.StatusFromError(err)

		c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// ->  get accounts withdrawals of all user
func (r *AdminController) GetWithdrawals(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := r.service.GetWithdrawalsService(limit, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: fiber.Map{
			"data":  data,
			"total": total,
		},
	})
}

// -> get users deatails
func (a *AdminController) GetUsersDetail(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	data, total, err := a.service.GetUsersDetailService(limit, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			constants.Error: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: fiber.Map{
			"data":  data,
			"total": total,
		},
	})
}

// -> block user

func (a *AdminController) BlockUser(c *fiber.Ctx) error {
	UserID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := a.service.BlockUserService(UserID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> unblock user
func (a *AdminController) UnBlockUser(c *fiber.Ctx) error {
	UserID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := a.service.UnBlockUserService(UserID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> get all flaged posts
func (a *AdminController) GetAllFlagedPost(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	offset, err := strconv.Atoi(c.Query("offset", "0"))

	Data, err := a.service.GetAllFlagedPostService(limit, offset)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> get all wallets details
func (s *AdminController) GetWallets(c *fiber.Ctx) error {

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	offset, err := strconv.Atoi(c.Query("offset", "0"))

	Data, err := s.service.GetWalletsService(limit, offset)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> block wallet
func (s *AdminController) BlockWallet(c *fiber.Ctx) error {
	WalletID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := s.service.BlockWalletService(WalletID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> unblock wallet
func (s *AdminController) UnBlockWallet(c *fiber.Ctx) error {
	WalletID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := s.service.UnBlockWalletService(WalletID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// ->verify Bank account
func (s *AdminController) VerifyAccount(c *fiber.Ctx) error {
	BankID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := s.service.VerifyAccountService(BankID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> remove the flaged post
func (s *AdminController) RemoveTheFlaggedPost(c *fiber.Ctx) error {
	PostID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := s.service.RemoveTheFlaggedPostService(PostID); err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: fiber.Map{
			"status":  constants.SUCESSCODE,
			"Message": "deleted succesfully",
			"PostID":  PostID,
		},
	})
}

// -> mark as the post safe
func (a *AdminController) SafePosting(c *fiber.Ctx) error {
	PostID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := a.service.SafePostingService(PostID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})

}

// -> get admin profile page
func (a *AdminController) GetProfile(c *fiber.Ctx) error {
	AdminID, _ := c.Locals("user_id").(uint)

	Data, err := a.service.GetProfile(AdminID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})

}

// -> upadate admin profile page
func (s *AdminController) UpadteProfile(c *fiber.Ctx) error {
	var Inputs domain.UpdateProfile
	AdminID, _ := c.Locals("user_id").(uint)

	if err := c.BodyParser(&Inputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	Data, err := s.service.UpdateProfileService(Inputs, AdminID)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(Data),
	})
}

// -> update user profile page
func (s *AdminController) UpdateUserProfile(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var Inputs domain.UpdateProfile

	if err := c.BodyParser(&Inputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	data, err := s.service.UpdateUserProfileService(Inputs, uint(id))

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse(data),
	})
}

// -> admin add user or any role person
func (s *AdminController) AddUser(c *fiber.Ctx) error {
	var Inputs domain.UpdateProfile

	if err := c.BodyParser(&Inputs); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	err := s.service.AddUserService(Inputs)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("User craeted succesfully"),
	})
}

// -> delete that user
func (a *AdminController) DeleteUser(c *fiber.Ctx) error {
	UserId, _ := c.ParamsInt("id")

	err := a.service.DeleteUserService(uint(UserId))

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("User Deleted succesfully"),
	})
}

// -> add main intrest
func (s *AdminController) AddMainIntrest(c *fiber.Ctx) error {
	var Mainintrest domain.MainCategory

	if err := c.BodyParser(&Mainintrest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	err := s.service.AddMainIntrestService(Mainintrest)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("main intrest added"),
	})
}

// -> upadte main intrest
func (a *AdminController) UpadteMainIntrest(c *fiber.Ctx) error {
	MainID, _ := c.ParamsInt("id")
	var Mainintrest domain.MainCategory

	if err := c.BodyParser(&Mainintrest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	err := a.service.UpadteMainIntrestService(Mainintrest, uint(MainID))

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("Updated main intrest"),
	})
}

// -> add sub intrest
func (a *AdminController) AddSubIntrest(c *fiber.Ctx) error {
	var SubIntrest domain.CreateSubInterestRequest

	if err := c.BodyParser(&SubIntrest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	err := a.service.AddSubIntrestService(SubIntrest)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("added sub intrest"),
	})
}

// -> upadte sub intrest
func (a *AdminController) UpdateSubIntrest(c *fiber.Ctx) error {
	SubID, _ := c.ParamsInt("id")
	var SubIntrest domain.SubInterest

	if err := c.BodyParser(&SubIntrest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	err := a.service.UpdateSubIntrestService(SubIntrest, uint(SubID))

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted sub intrest"),
	})
}

// -> delete the main intrest
func (a *AdminController) DeleteIntrest(c *fiber.Ctx) error {
	Mainid, _ := c.ParamsInt("id")

	err := a.service.DeleteIntrestService(Mainid)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted sub intrest"),
	})
}

// -> delete the Sub intrest
func (a *AdminController) DeleteSubIntrest(c *fiber.Ctx) error {
	Subid, _ := c.ParamsInt("id")

	err := a.service.DeleteSubIntrestService(Subid)

	if err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted sub intrest"),
	})
}

// -> BLOCK THE POST
func (a *AdminController) BlockPost(c *fiber.Ctx) error {
	PostID, _ := c.ParamsInt("id")

	if err := a.service.BlockPostService(uint(PostID)); err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted sub intrest"),
	})
}

// -> UNBLOCK THE POST
func (a *AdminController) UnBlockPost(c *fiber.Ctx) error {
	PostID, _ := c.ParamsInt("id")

	if err := a.service.UnBlockPostService(uint(PostID)); err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted sub intrest"),
	})
}

//-> update User post intrests

func (a *AdminController) UpdatePostIntrests(c *fiber.Ctx) error {
	PostID, _ := c.ParamsInt("id")
	var Request domain.UpdatePostInterestRequest

	if err := c.BodyParser(&Request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(constants.BADREQUEST, err),
		})
	}

	if err := a.service.EditInrtestService(uint(PostID), Request); err != nil {
		status := response.StatusFromError(err)

		return c.Status(status).JSON(fiber.Map{
			constants.Error: response.ErrorMessage(status, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.Sucess: response.SuccessResponse("upadted intrest of posts"),
	})
}
