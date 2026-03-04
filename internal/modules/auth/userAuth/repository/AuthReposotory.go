package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain"
)

// → Postgres implementation
type AuthRespository struct {
	DB    *gorm.DB
	redis *redis.Client
}

func NewPostgresAuthRepo(db *gorm.DB, rds *redis.Client) domain.AuthRepo {
	return &AuthRespository{DB: db, redis: rds}
}

func (r *AuthRespository) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}

func (r *AuthRespository) FindAnything(model interface{}, Query, Any interface{}) error {
	return r.DB.Where(Query, Any).First(model).Error
}

func (r *AuthRespository) FindAnythingProload(model interface{}, Query, Any interface{}) error {
	return r.DB.
		Preload("Permissions").
		Where(Query, Any).
		First(model).Error
}
