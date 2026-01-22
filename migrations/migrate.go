package migrations

import "gorm.io/gorm"

func Migrations(db *gorm.DB) {
	db.AutoMigrate()
}
