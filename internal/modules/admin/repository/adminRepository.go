package repository

import (
	"thinkdrop-backend/internal/modules/admin/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) domain.AdminRepo {
	return &Repository{DB: db}
}
