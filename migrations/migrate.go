package migrations

import (
	"fmt"
	"log"
	AuthDomain "thinkdrop-backend/internal/modules/auth/userAuth/domain"
	InterstDomain "thinkdrop-backend/internal/modules/interest/domain"

	"gorm.io/gorm"
)

// -> enitre migrations happens here
func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		AuthDomain.User{},
		InterstDomain.MainInterest{},
		InterstDomain.SubInterest{},
	)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	fmt.Print("Migration success✅")
}
