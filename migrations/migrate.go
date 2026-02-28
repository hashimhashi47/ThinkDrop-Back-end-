package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	domain "thinkdrop-backend/internal/Common"
)

// -> enitre migrations happens here
func Migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		domain.User{},
		domain.MainInterest{},
		domain.SubInterest{},
		domain.Post{},
		domain.Report{},
		domain.Comment{},
		domain.UserFollow{},
		domain.Avatar{},
		domain.Wallet{},
		domain.BankAccount{},
		domain.Like{},
		domain.Withdrawal{},
		domain.Message{},
		domain.Conversation{},
		domain.ReportComplaints{},
	)
	if err != nil {
		log.Fatal("Migration error", err)
	}
	fmt.Print("Migration success✅")
}
