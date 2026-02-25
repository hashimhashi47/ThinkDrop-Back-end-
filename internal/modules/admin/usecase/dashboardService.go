package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	DomainAdmin "thinkdrop-backend/internal/modules/admin/domain"
	"thinkdrop-backend/internal/modules/admin/websocket"
)

type AdminService struct {
	repo DomainAdmin.AdminRepo
	hub  *websocket.Hub
}

func NewAdminService(r DomainAdmin.AdminRepo, h *websocket.Hub) *AdminService {
	return &AdminService{repo: r, hub: h}
}

// -> get tha statiscs for admin dashboard logic
func (a *AdminService) GetdashboardDetailsService() (interface{}, error) {
	var err error
	var Wallet domain.Wallet
	var User domain.User
	var dashboard DomainAdmin.DashboardResponse

	UserCount, err := a.repo.GetCounts("users")
	postsCount, err := a.repo.GetCounts("posts")
	reportsCount, err := a.repo.GetCounts("reports")

	if err != nil {
		return nil, errors.New("failed to get the statics")
	}

	if err := a.repo.FindTopUser(&Wallet); err != nil {
		return nil, errors.New("failed to get wallet details")
	}

	if err := a.repo.FindbyArgs(&User, "id = ?", Wallet.UserID); err != nil {
		return nil, errors.New("failed to get user details")
	}

	dashboard = DomainAdmin.DashboardResponse{
		TotalUsers:   UserCount,
		ReportsCount: reportsCount,
		TotalPosts:   postsCount,
		TopRedeemer: DomainAdmin.TopRedeemer{
			Name:   User.FullName,
			Points: Wallet.PointsAvailable,
		},
	}

	return dashboard, nil
}

// -> get all users post and details logics
func (a *AdminService) GetAllusersPostService(limit, offset int) (interface{}, int64, error) {
	var Post []domain.Post

	var PostResponse []domain.PostFeedResponse

	if err := a.repo.FindAll(&Post, limit, offset, "User", "SubInterests"); err != nil {
		return nil, 0, errors.New("failed to find the posts")
	}

	total, err := a.repo.Count(&domain.Post{})
	if err != nil {
		return nil, 0, err
	}

	for _, v := range Post {
		// collect interests
		interests := make([]domain.PostInterestDTO, 0, len(v.SubInterests))
		for _, si := range v.SubInterests {
			interests = append(interests, domain.PostInterestDTO{
				PID:  si.ID,
				Name: si.Name,
			})
		}
		PostResponse = append(PostResponse, domain.PostFeedResponse{
			ID:          v.ID,
			Content:     v.Content,
			CreatedAt:   v.CreatedAt,
			LikeCount:   v.LikeCount,
			Interests:   interests,
			IsBlocked:   v.Blocked,
			ReportCount: v.ReportCount,
			User: domain.PostUserDTO{
				UID:           v.UserID,
				Name:          v.User.FullName,
				AnonymousName: v.User.AnonymousName,
				ImageURL:      v.User.ImageURL,
			},
		})
	}
	return PostResponse, total, nil
}
