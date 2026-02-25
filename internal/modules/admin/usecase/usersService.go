package usecase

import (
	"errors"
	"fmt"
	AdminDomain "thinkdrop-backend/internal/modules/admin/domain"
	hashpass "thinkdrop-backend/pkg/hashPass"
	"time"

	domain "thinkdrop-backend/internal/Common"
)

func (a *AdminService) GetUsersDetailService(limit, offset int) (interface{}, int64, error) {

	var users []domain.User

	total, err := a.repo.Count(&domain.User{})
	if err != nil {
		return nil, 0, errors.New("failed to count transactions")
	}

	if err := a.repo.FindWithoutPreload(&users, limit, offset); err != nil {
		return nil, 0, errors.New("failed to find users")
	}

	var response []AdminDomain.AdminUserDTO
	for _, v := range users {
		// if v.Role == constants.User {
		response = append(response, AdminDomain.AdminUserDTO{
			ID:        v.ID,
			FullName:  v.FullName,
			Email:     v.Email,
			Role:      v.Role,
			IsBlocked: v.IsBlocked,
			Verify:    v.Verify,
			ImageURL:  v.ImageURL,
			CreatedAt: v.CreatedAt,
		})
		// }
	}

	return response, total, nil
}

// -> block user logic
func (a *AdminService) BlockUserService(UserID int) (interface{}, error) {
	var User domain.User

	if err := a.repo.UpdateColumn(&User, "id = ?", UserID, "IsBlocked", true); err != nil {
		return nil, errors.New("failed to block the user")
	}

	return User, nil
}

// -> unblock user logic
func (a *AdminService) UnBlockUserService(UserID int) (interface{}, error) {
	var User domain.User

	if err := a.repo.UpdateColumn(&User, "id = ?", UserID, "IsBlocked", false); err != nil {
		return nil, errors.New("failed to block the user")
	}

	return User, nil
}

func (a *AdminService) UpdateUserProfileService(Inputs AdminDomain.UpdateProfile,
	UserID uint) (interface{}, error) {

	var err error

	if Inputs.FullName != "" {
		err = a.repo.UpdateColumn(&domain.User{}, "id = ?", UserID, "FullName", Inputs.FullName)
	}

	if Inputs.Email != "" {
		err = a.repo.UpdateColumn(&domain.User{}, "id = ?", UserID, "Email", Inputs.Email)
	}

	if Inputs.Password != "" {
		HashedPass, _ := hashpass.GenerateHashedPassword(Inputs.Password)
		err = a.repo.UpdateColumn(&domain.User{}, "id = ?", UserID, "Password", HashedPass)
	}

	if Inputs.Role != "" {
		err = a.repo.UpdateColumn(&domain.User{}, "id = ?", UserID, "Role", Inputs.Role)
	}

	if err != nil {
		return nil, err
	}

	var user domain.User
	if err := a.repo.FindbyArgs(&user, "id = ?", UserID); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AdminService) AddUserService(inputs AdminDomain.UpdateProfile) error {

	hashpass, _ := hashpass.GenerateHashedPassword(inputs.Password)

	user := domain.User{
		FullName:      inputs.FullName,
		Email:         inputs.Email,
		Password:      hashpass,
		Role:          inputs.Role,
		AnonymousName: fmt.Sprintf("anon_%d", time.Now().UnixNano()),
	}

	if err := a.repo.Insert(&user); err != nil {
		return errors.New("failed to add this user")
	}

	return nil
}

// -> delete the user from database
func (a *AdminService) DeleteUserService(UserID uint) error {

	if err := a.repo.Delete(&domain.User{}, "id = ?", UserID); err != nil {
		return errors.New("failed to delete the user")
	}
	return nil
}
