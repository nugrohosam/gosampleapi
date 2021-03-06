package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// GetUserRole ..
func GetUserRole(search, perPage, page, order string) (userRoleRepo.UserRoles, int, error) {

	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	userRoles, total, err := userRoleRepo.Get(search, limit, offset, orderBy)
	if err != nil {
		return nil, 0, err
	}

	return userRoles, total, nil
}

// ShowUserRole ..
func ShowUserRole(ID int) userRoleRepo.UserRole {
	userRole := userRoleRepo.Find(ID)
	return userRole
}

// GetUserRoleWithUserID ..
func GetUserRoleWithUserID(userID int) userRoleRepo.UserRoles {
	userRoles := userRoleRepo.GetByUserID(userID)
	return userRoles
}

// IsHaveAnyRole ...
func IsHaveAnyRole(userID int, roleName []string) (bool, error) {
	isExist := userRoleRepo.IsExistsByUserIDAndRoleName(userID, roleName)
	return isExist, nil
}

// CreateUserRole ...
func CreateUserRole(userID, roleID int) error {
	_, err := userRoleRepo.Create(userID, roleID)
	return err
}

// UpdateUserRole ...
func UpdateUserRole(ID, roleID, userID int) error {
	_, err := userRoleRepo.Update(ID, roleID, userID)
	return err
}

// DeleteUserRole ...
func DeleteUserRole(ID int) error {
	_, err := userRoleRepo.Delete(ID)
	return err
}
