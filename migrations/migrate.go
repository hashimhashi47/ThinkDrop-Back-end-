package migrations

import (
	"fmt"
	"log"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain"

	"gorm.io/gorm"
)

// -> enitre migrations happens here
func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		domain.User{},
	)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	fmt.Print("Migration success✅")
}
