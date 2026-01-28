package databaserepository

import (
	"gorm.io/gorm"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/repository/authRepository"
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

func (r *PostgresRepo) FindAnything(model interface{}, Query, Any string) error {
	return r.DB.Where(Query, Any).First(model).Error
}
