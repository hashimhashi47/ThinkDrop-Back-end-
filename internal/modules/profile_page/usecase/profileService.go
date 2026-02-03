package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	ProfileDomain "thinkdrop-backend/internal/modules/profile_page/domain"
)

type ProfileService struct {
	repo ProfileDomain.ProfileRepo
}

func NewProfileService(r ProfileDomain.ProfileRepo) *ProfileService {
	return &ProfileService{repo: r}
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
	}

	return Response, nil
}

// -> swow other users profile details
func (r *ProfileService) ShowOtherUserProfileService(id int) (domain.ProfileResponseDTO, error) {
	var UserDetails domain.User

	if err := r.repo.Find(&UserDetails, "id = ?", id, "Posts.SubInterest"); err != nil {
		return domain.ProfileResponseDTO{}, errors.New("failed to find the user")
	}

	writings := make([]domain.WritingDTO, 0, len(UserDetails.Posts))
	for _, v := range UserDetails.Posts {
		writings = append(writings, domain.WritingDTO{
			ID:        v.ID,
			Content:   v.Content,
			Intrests:  v.SubInterest.Name,
			CreatedAt: v.CreatedAt,
		})
	}

	return domain.ProfileResponseDTO{
		AnonymousName: UserDetails.AnonymousName,
		WritingsCount: len(UserDetails.Posts),
		Bio:           UserDetails.Bio,
		Writings:      writings,
	}, nil
}

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
