package main

import (
	"fmt"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/pkg/constants"
	hashpass "thinkdrop-backend/pkg/hashPass"
)

func main() {

	database.Connection()
	password := "admin123"
	db := database.DB

	hashed, _ := hashpass.GenerateHashedPassword(password)
	var existing domain.User
	err := db.Where("email = ?", "admin@gmail.com").First(&existing).Error

	if err != nil {
		admin := domain.User{
			FullName: "admin",
			Verify:   true,
			Email:    "admin@gmail.com",
			Password: hashed,
			Role:     constants.Admin,
		}
		db.Create(&admin)
		fmt.Println("Admin created ✅")
	} else {
		fmt.Println("Admin already exists ⚡")
	}
}
