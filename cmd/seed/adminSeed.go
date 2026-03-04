package seed

import (
	"errors"
	"fmt"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/pkg/constants"
	hashpass "thinkdrop-backend/pkg/hashPass"

	"gorm.io/gorm"
)

func CreateAdmin(db *gorm.DB) {
	password := "admin123"
	hashed, err := hashpass.GenerateHashedPassword(password)
	if err != nil {
		panic("Failed to hash password")
	}

	var existing domain.User
	err = db.Where("email = ?", "admin@gmail.com").First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		admin := domain.User{
			FullName: "admin",
			Verify:   true,
			Email:    "admin@gmail.com",
			Password: hashed,
			Role:     constants.Admin,
		}

		if err := db.Create(&admin).Error; err != nil {
			panic("Failed to create admin")
		}

		fmt.Println("Admin created ✅")
	} else if err == nil {
		fmt.Println("Admin already exists ⚡")
	} else {
		panic(err)
	}
}
