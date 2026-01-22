package repository

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/repository"

	"gorm.io/gorm"
)

// → Postgres implementation

type PostgresRepo struct {
	DB *gorm.DB
}

func NewPostgresAuthRepo(db *gorm.DB) repository.AuthRespository {
	return &PostgresRepo{DB: db}
}

func (r *PostgresRepo) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}
