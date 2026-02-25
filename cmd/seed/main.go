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

	admin := domain.User{
		FullName: "admin",
		Verify:   true,
		Email:    "admin@gmail.com",
		Password: hashed,
		Role:     constants.Admin,
	}

	db.Create(&admin)
	fmt.Println("worked✅")
}
