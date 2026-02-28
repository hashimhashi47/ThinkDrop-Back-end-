package usecase

import (
	"errors"
	"strconv"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/internal/modules/admin/usecase"
	PostDomain "thinkdrop-backend/internal/modules/post/domain"

	"gorm.io/gorm"
)

type PostService struct {
	repo         PostDomain.PostRepo
	adminService usecase.AdminService
}

func NewPostService(r PostDomain.PostRepo, aS usecase.AdminService) *PostService {
	return &PostService{repo: r, adminService: aS}
}

// -> user can upload their post it also have check limit request
func (r *PostService) AddPostService(Post domain.CreatePostRequest, UserID uint) (domain.Post, error) {
	var User domain.User
	var AddPost domain.Post
	var subInterests []domain.SubInterest

	if err := r.repo.FindAnything(&User, "id = ?", UserID); err != nil {
		return domain.Post{}, errors.New("failed to find user")
	}

	ok, err := r.repo.AllowPost(UserID)

	if err != nil {
		return domain.Post{}, errors.New("Too many request")
	}

	if !ok {
		return domain.Post{}, errors.New("try again after 30 seconds")
	}

	if Post.InterestIDs != nil {
		if err := r.repo.FindAll(&subInterests, "id IN ?", Post.InterestIDs); err != nil {
			return domain.Post{}, errors.New("failed to fetch interests")
		}

		if len(subInterests) != len(Post.InterestIDs) {
			return domain.Post{}, errors.New("invalid interest ids")
		}
	}

	subInterests = append(subInterests, domain.SubInterest{
		ID: 1,
	})

	AddPost = domain.Post{
		Content:      Post.Content,
		UserID:       UserID,
		SubInterests: subInterests,
	}

	if err := r.repo.Insert(&AddPost); err != nil {
		return domain.Post{}, errors.New("Failed to post your Blog")
	}

	dashboardData, _ := r.adminService.GetdashboardDetailsService()
	PostData, _, _ := r.adminService.GetAllusersPostService(10, 0)
	r.adminService.Broadcast("posts", "UPDATE_POST", PostData)
	r.adminService.Broadcast("dashboard", "UPDATE", dashboardData)

	return AddPost, nil
}

// -> show All users post can show their acccounts
func (r *PostService) ShowPostsServices(userID uint) ([]domain.PostResponse, error) {
	var posts []domain.Post

	if err := r.repo.FindAnyWithpreload(&posts, "user_id = ?", userID, "SubInterests"); err != nil {
		return nil, err
	}

	var response []domain.PostResponse

	for _, post := range posts {
		var interests []string

		for _, v := range post.SubInterests {
			interests = append(interests, v.Name)
		}
		response = append(response, domain.PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			SubInterest:  interests,
		})
	}

	return response, nil
}

// -> user feed sesion it will customised by the intrest
func (r *PostService) UserFeedService(userID uint, limit, offset int) ([]domain.PostFeedResponse, error) {
	var User domain.User
	var AllPosts []domain.Post

	if err := r.repo.FindByUser(&User, "id = ?", userID); err != nil {
		return nil, err
	}

	subIDs := make([]uint, 0, len(User.SelectedSubs))
	for _, sub := range User.SelectedSubs {
		subIDs = append(subIDs, sub.ID)
	}

	if len(subIDs) == 0 {
		return []domain.PostFeedResponse{}, nil
	}

	if err := r.repo.FindFeedPosts(&AllPosts, subIDs, limit, offset); err != nil {
		return nil, err
	}

	feed := make([]domain.PostFeedResponse, 0, len(AllPosts))

	for _, post := range AllPosts {
		if !post.Blocked {
			interests := make([]domain.PostInterestDTO, 0)

			for _, si := range post.SubInterests {
				interests = append(interests, domain.PostInterestDTO{
					PID:  si.ID,
					Name: si.Name,
				})
			}

			ok, _ := r.repo.FindLikeByUserID(User.ID, post.ID)

			feed = append(feed, domain.PostFeedResponse{
				ID:            post.ID,
				Content:       post.Content,
				CreatedAt:     post.CreatedAt,
				LikeCount:     post.LikeCount,
				Interests:     interests,
				IsUserIsLiked: ok,
				User: domain.PostUserDTO{
					UID:           post.UserID,
					AnonymousName: post.User.AnonymousName,
					ImageURL:      post.User.ImageURL,
				},
			})
		}
	}
	return feed, nil
}

// -> add like on post logic
func (r *PostService) LikePostService(UserID uint, PostID int) (bool, error) {
	Like := domain.Like{
		UserID: UserID,
		PostID: uint(PostID),
	}

	var post domain.Post

	if err := r.repo.FindAnything(&post, "id = ?", PostID); err != nil {
		return false, errors.New("failed to find the post")
	}

	if err := r.repo.Create(&Like); err != nil {
		return false, errors.New("already liked the post")
	}

	if err := r.repo.UpdateUserWalletByPostID(post.ID, 2); err != nil {
		return false, errors.New("failed to update wallet points")
	}

	if err := r.repo.UpdateColumn(&domain.Post{}, "id = ?", PostID, "like_count",
		gorm.Expr("like_count + 1")); err != nil {
		return false, err
	}

	likedata, _, _ := r.adminService.GetWithdrawalsService(10, 0)
	walletdata, _ := r.adminService.GetWalletsService(10, 0)
	r.adminService.Broadcast("accounts", "UPDATE_ACCOUNT", likedata)
	r.adminService.Broadcast("wallets", "UPDATE_WALLET", walletdata)
	return true, nil
}

// -> unlike post
func (r *PostService) UnLikePostService(UserID uint, PostID int) (bool, error) {
	result, err := r.repo.DeleteWhere(&domain.Like{}, "user_id = ? AND post_id = ?", UserID, PostID)
	if err != nil {
		return false, err
	}

	var post domain.Post

	if err := r.repo.FindAnything(&post, "id = ?", PostID); err != nil {
		return false, errors.New("failed to find the post")
	}

	if result.RowsAffected == 0 {
		return false, errors.New("post not liked")
	}

	if err := r.repo.UpdateColumn(&domain.Post{}, "id = ?", PostID, "like_count",
		gorm.Expr("like_count - 1")); err != nil {
		return false, err
	}

	if err := r.repo.UpdateUserWalletByPostID(post.ID, -2); err != nil {
		return false, errors.New("failed to update wallet points")
	}

	likedata, _, _ := r.adminService.GetWithdrawalsService(10, 0)
	walletdata, _ := r.adminService.GetWalletsService(10, 0)
	r.adminService.Broadcast("accounts", "UPDATE_ACCOUNT", likedata)
	r.adminService.Broadcast("wallets", "UPDATE_WALLET", walletdata)
	return true, nil
}

// -> Report post logics

func (s *PostService) ReportPostService(data domain.ReportPostRequest, UserID uint) (interface{}, error) {
	var Post domain.Post
	var Report domain.Report

	isOk, err := s.repo.ReportRateLimit(strconv.FormatUint(uint64(data.PostID), 10))

	if err != nil {
		return "", err
	}

	if !isOk {
		return "", errors.New("Request limit exceeded, wait for report again")
	}

	if err := s.repo.FindAnyWithpreload(&Post, "id = ?", data.PostID, "Reports"); err != nil {
		return nil, err
	}

	Report = domain.Report{
		PostID:      data.PostID,
		UserID:      UserID,
		Reason:      data.Reason,
		Description: data.Description,
	}

	Post.Reports = append(Post.Reports, Report)

	if err := s.repo.Create(&Report); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateColumn(&domain.Post{}, "id = ?", data.PostID, "report_count",
		gorm.Expr("report_count + ?", 1)); err != nil {
		return nil, err
	}

	// -> wesocket update
	dashboardData, _ := s.adminService.GetdashboardDetailsService()
	PostData, _, _ := s.adminService.GetAllusersPostService(10, 0)
	ReportData, _ := s.adminService.GetAllFlagedPostService(10, 0)
	s.adminService.Broadcast("dashboard", "UPDATE", dashboardData)
	s.adminService.Broadcast("posts", "UPDATE_POST", PostData)
	s.adminService.Broadcast("reports", "UPDATE_REPORT", ReportData)
	return Report, nil
}
