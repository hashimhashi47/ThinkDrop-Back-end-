package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	AdminDomain "thinkdrop-backend/internal/modules/admin/domain"
	hashpass "thinkdrop-backend/pkg/hashPass"
)

func (a *AdminService) GetProfile(AdminID uint) (interface{}, error) {
	var Admin domain.User

	if err := a.repo.FindbyArgs(&Admin, "id = ?", AdminID); err != nil {
		return nil, errors.New("Failed to find the admin details")
	}

	Data := AdminDomain.AdminProfile{
		ID:        AdminID,
		Role:      Admin.Role,
		Email:     Admin.Email,
		Name:      Admin.FullName,
		CreatedAt: Admin.CreatedAt,
		ImageURL:  Admin.ImageURL,
	}
	return Data, nil
}

func (a *AdminService) UpdateProfileService(Input AdminDomain.UpdateProfile, AdminID uint) (interface{}, error) {
	var Admin domain.User
	var err error

	var response AdminDomain.AdminProfile

	if err := a.repo.FindbyArgs(&Admin, "id = ?", AdminID); err != nil {
		return nil, errors.New("Failed to find the admin details")
	}

	if Input.Email != "" {
		err = a.repo.UpdateColumn(&Admin, "id = ?", AdminID, "Email", Input.Email)
		response.Email = Input.Email
	}

	if Input.FullName != "" {
		err = a.repo.UpdateColumn(&Admin, "id = ?", AdminID, "FullName", Input.FullName)
		response.Name = Input.FullName
	}

	if Input.Password != "" {
		hashed, _ := hashpass.GenerateHashedPassword(Input.Password)
		err = a.repo.UpdateColumn(&Admin, "id = ?", AdminID, "Password", hashed)
	}

	if Input.ImageURL != "" {
		err = a.repo.UpdateColumn(&Admin, "id = ?", AdminID, "ImageURL", Input.ImageURL)
		response.ImageURL = Input.ImageURL
	}

	if err != nil {
		return nil, errors.New("failed to update profile")
	}

	return response, nil
}


