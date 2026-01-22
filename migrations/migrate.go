package migrations

import (
	"fmt"
	"log"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/entity"

	"gorm.io/gorm"
)

// -> enitre migrations happens here
func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.User{},
	)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	fmt.Print("Migration success✅")
}
