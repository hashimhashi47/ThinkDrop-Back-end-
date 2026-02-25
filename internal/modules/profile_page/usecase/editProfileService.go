package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
)

// -> get avatars logic
func (r *ProfileService) GetAvatarsService() ([]domain.Avatar, error) {
	var Avatars []domain.Avatar

	if err := r.repo.FindAll(&Avatars); err != nil {
		return nil, errors.New("failed to get the data")
	}

	return Avatars, nil
}

// -> Edit profile logics
func (r *ProfileService) EditProfileService(UserID uint, UserInputs domain.EditProfile) (domain.UserProfileResponse, error) {
	var User domain.User

	if err := r.repo.Find(&User, "id = ?", UserID); err != nil {
		return domain.UserProfileResponse{}, errors.New("failed to find the user")
	}

	if UserInputs.AnonymousName != "" {
		User.AnonymousName = UserInputs.AnonymousName
	}

	if UserInputs.Bio != "" {
		User.Bio = UserInputs.Bio
	}

	if UserInputs.ImageURL != "" {
		User.ImageURL = UserInputs.ImageURL
	}

	if err := r.repo.Save(&User); err != nil {
		return domain.UserProfileResponse{}, errors.New("failed to update the data")
	}

	response := domain.UserProfileResponse{
		AnonymousName: User.AnonymousName,
		Bio:           User.Bio,
		ProfileAvatar: User.ImageURL,
	}

	Userdata, _, _ := r.AdminService.GetUsersDetailService(10, 0)
	r.AdminService.Broadcast("users", "UPDATE_USER", Userdata)
	return response, nil
}
