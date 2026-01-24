package databaserepository

import (
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/repository/authRepository"

	"gorm.io/gorm"
)

// → Postgres implementation

type PostgresRepo struct {
	DB *gorm.DB
}

func NewPostgresAuthRepo(db *gorm.DB) authrepository.AuthRespository {
	return &PostgresRepo{DB: db}
}

func (r *PostgresRepo) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}
