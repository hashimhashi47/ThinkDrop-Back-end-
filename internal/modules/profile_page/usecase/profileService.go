package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	AdminUsecase "thinkdrop-backend/internal/modules/admin/usecase"
	ProfileDomain "thinkdrop-backend/internal/modules/profile_page/domain"
)

type ProfileService struct {
	repo         ProfileDomain.ProfileRepo
	AdminService AdminUsecase.AdminService
}

func NewProfileService(r ProfileDomain.ProfileRepo, as AdminUsecase.AdminService) *ProfileService {
	return &ProfileService{repo: r, AdminService: as}
}

// -> show the own user details
func (r *ProfileService) ShowProfileService(UserID uint) (domain.ProfilePage, error) {
	var User domain.User

	if err := r.repo.Find(&User, "id = ?", UserID, "Following", "Followers", "Posts"); err != nil {
		return domain.ProfilePage{}, errors.New("failed find the user")
	}

	Response := domain.ProfilePage{
		AnonymousName:  User.AnonymousName,
		FollowersCount: len(User.Followers),
		FollowingCount: len(User.Following),
		WritingsCount:  len(User.Posts),
		Bio:            User.Bio,
		Avatar:         User.ImageURL,
	}

	return Response, nil
}

// -> show other user's profile details + follow status
func (r *ProfileService) ShowOtherUserProfileService(profileUserID int, currentUserID uint) (domain.ProfileResponseDTO, error) {

	var UserDetails domain.User

	if err := r.repo.Find(&UserDetails, "id = ?", profileUserID, "Posts.SubInterests", "Followers"); err != nil {
		return domain.ProfileResponseDTO{}, errors.New("failed to find the user")
	}

	isFollowing := false
	for _, f := range UserDetails.Followers {
		if f.FollowerID == currentUserID {
			isFollowing = true
			break
		}
	}

	writings := make([]domain.WritingDTO, 0, len(UserDetails.Posts))
	for _, v := range UserDetails.Posts {

		var interests []string

		for _, v := range v.SubInterests {
			interests = append(interests, v.Name)
		}

		writings = append(writings, domain.WritingDTO{
			ID:        v.ID,
			Content:   v.Content,
			Intrests:  interests,
			ImageURL:  UserDetails.ImageURL,
			CreatedAt: v.CreatedAt,
		})
	}

	return domain.ProfileResponseDTO{
		AnonymousName: UserDetails.AnonymousName,
		WritingsCount: len(UserDetails.Posts),
		Bio:           UserDetails.Bio,
		ImageURL:      UserDetails.ImageURL,
		Writings:      writings,
		IsFollowing:   isFollowing,
	}, nil
}

// -> follow a user service
func (r *ProfileService) FollowUserService(UserID uint,
	OtherUserID int) (domain.UserProfileResponse, domain.UserProfileResponse, error) {

	count, err := r.repo.CountFollow(UserID, uint(OtherUserID))

	if err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	if count > 0 {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, errors.New("already following")
	}

	follow := domain.UserFollow{
		FollowerID: UserID,
		FollowedID: uint(OtherUserID),
	}

	if err := r.repo.Create(&follow); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{},
			errors.New("failed to following")
	}

	var User, OtherUser domain.User

	if err := r.repo.Find(&User, "id = ?", UserID, "Following.Followed"); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	if err := r.repo.Find(&OtherUser, "id = ?", OtherUserID, "Followers.Follower"); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	return domain.MapUserToProfile(User), domain.MapUserToProfile(OtherUser), nil
}

// -> unfollow logic
func (r *ProfileService) UnfollowUser(userID uint,
	otherUserID uint) (domain.UserProfileResponse, domain.UserProfileResponse, error) {

	if err := r.repo.Unfollow(userID, otherUserID); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	var User, OtherUser domain.User
	if err := r.repo.Find(&User, "id = ?", userID, "Following", "Followers"); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	if err := r.repo.Find(&OtherUser, "id = ?", otherUserID, "Following", "Followers"); err != nil {
		return domain.UserProfileResponse{}, domain.UserProfileResponse{}, err
	}

	return domain.MapUserToProfile(User), domain.MapUserToProfile(OtherUser), nil
}

// -> get all Users writings logic
func (r *ProfileService) GetAllWritingsService(UserID uint, limit, offset int) ([]domain.PostFeedResponse, error) {

	posts, err := r.repo.GetUserPosts(UserID, limit, offset)
	if err != nil {
		return nil, errors.New("failed to get writings")
	}

	var Posts []domain.PostFeedResponse

	for _, v := range posts {
		if !v.Blocked {
			// collect interests
			interests := make([]domain.PostInterestDTO, 0, len(v.SubInterests))
			for _, si := range v.SubInterests {
				interests = append(interests, domain.PostInterestDTO{
					PID:  si.ID,
					Name: si.Name,
				})
			}
			Posts = append(Posts, domain.PostFeedResponse{
				ID:        v.ID,
				Content:   v.Content,
				CreatedAt: v.CreatedAt,
				LikeCount: v.LikeCount,
				Interests: interests,
				User: domain.PostUserDTO{
					UID:           v.UserID,
					Name:          v.User.FullName,
					AnonymousName: v.User.AnonymousName,
					ImageURL:      v.User.ImageURL,
				},
			})
		}
	}

	return Posts, nil
}

// -> get all users followers logic
func (r *ProfileService) GetAllFollowersService(UserID interface{}) ([]domain.FollowUserResponse, error) {
	var User domain.User

	var Followers []domain.FollowUserResponse

	if err := r.repo.Find(&User, "id = ?", UserID, "Followers.Follower", "Following.Followed"); err != nil {
		return nil, errors.New("failed to get the Followers")
	}

	following := make(map[uint]bool)

	for _, v := range User.Following {
		following[v.FollowedID] = true
	}

	for _, c := range User.Followers {

		isFollowing := following[c.FollowerID]

		Followers = append(Followers, domain.FollowUserResponse{
			UserID:        c.Follower.ID,
			AnonymousName: c.Follower.AnonymousName,
			ImageURL:      c.Follower.ImageURL,
			IsFollower:    true,
			IsFollowing:   isFollowing,
		})
	}

	return Followers, nil
}

// -> get all users following logic
func (r *ProfileService) GetAllFollowingService(UserID interface{}) ([]domain.FollowUserResponse, error) {
	var User domain.User

	var Following []domain.FollowUserResponse

	if err := r.repo.Find(&User, "id = ?", UserID, "Following.Followed", "Followers"); err != nil {
		return nil, errors.New("failed to get the Followers")
	}

	follower := make(map[uint]bool)

	for _, v := range User.Followers {
		follower[v.FollowerID] = true
	}

	for _, f := range User.Following {
		isFollower := follower[f.FollowedID]

		Following = append(Following, domain.FollowUserResponse{
			UserID:        f.Followed.ID,
			AnonymousName: f.Followed.AnonymousName,
			ImageURL:      f.Followed.ImageURL,
			IsFollowing:   true,
			IsFollower:    isFollower,
		})
	}

	return Following, nil
}

// -> get user intrests logic
func (r *ProfileService) GetUserIntrest(userID uint) (interface{}, error) {
	var user domain.User

	if err := r.repo.Find(&user, "id = ?", userID, "SelectedSubs"); err != nil {
		return nil, errors.New("failed to find the user")
	}

	return user.SelectedSubs, nil
}
