package seed

import (
	"fmt"
	"thinkdrop-backend/internal/modules/admin/domain"
	"thinkdrop-backend/pkg/constants"

	"gorm.io/gorm"
)

func CreatePermission(db *gorm.DB) error {

	permissions := []domain.Permission{
		{Slug: "dashboard", Name: "Dashboard"},
		{Slug: "users", Name: "Users"},
		{Slug: "accounts", Name: "Accounts"},
		{Slug: "wallet", Name: "Wallet"},
		{Slug: "interests", Name: "Interests"},
		{Slug: "surveillance", Name: "Surveillance"},
		{Slug: "all_posts", Name: "Posts"},
		{Slug: "reports", Name: "Reports"},
		{Slug: "roles_config", Name: "Config"},
		{Slug: "profile", Name: "Profile"},
	}

	for _, v := range permissions {
		if err := db.
			Where("slug = ?", v.Slug).
			FirstOrCreate(&v).Error; err != nil {
			return err
		}
	}

	fmt.Println("Permission seeding complete 🚀")
	return nil
}

func AdminPersmission(db *gorm.DB) error {
	var Admin domain.Role

	if err := db.Where("name = ?", constants.Admin).First(&Admin).Error; err != nil {
		return err
	}

	var Permisison []domain.Permission

	if err := db.Find(&Permisison).Error; err != nil {
		return err
	}

	return db.Model(&Admin).Association("Permissions").Append(&Permisison)
}
