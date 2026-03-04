package usecase

import (
	"errors"
	"fmt"
	"thinkdrop-backend/internal/modules/admin/domain"
)

func (a *AdminService) CreateRoleService(Input domain.CreateRoleRequest) error {
	var Role domain.Role
	Role.Name = Input.Name

	if err := a.repo.Insert(&Role); err != nil {
		return errors.New("Failed to creatate the role")
	}

	return nil
}

func (a *AdminService) GetRolesService() (interface{}, error) {
	var roles []domain.Role
	if err := a.repo.FindPreload(&roles, "Permissions"); err != nil {
		return nil, errors.New("failed to find the roles")
	}

	return roles, nil
}

func (a *AdminService) GetPermissionsService() (interface{}, error) {
	var Permisison []domain.Permission
	if err := a.repo.Find(&Permisison); err != nil {
		return nil, errors.New("failed to find the roles")
	}

	return Permisison, nil
}

func (a *AdminService) UpdateRolesService(id int, data domain.UpdateRoleInput) error {

	var role domain.Role
	fmt.Println("✅", id, data)

	if err := a.repo.FindbyArgs(&role, "id = ?", id); err != nil {
		return err
	}

	if data.Name != "" {
		if err := a.repo.UpdateColumn(&domain.Role{}, "id = ?", id, "name", data.Name); err != nil {
			return err
		}
	}

	if data.Permissions != nil {
		if err := a.repo.ReplaceRolePermissions(uint(id), data.Permissions); err != nil {
			return err
		}
	}

	return nil
}
