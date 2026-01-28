package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	"thinkdrop-backend/internal/modules/auth/userAuth/repository"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase"
)

// ->Init the auth service to pass the connections from database to controllers
func InitAuth(db *gorm.DB, rds *redis.Client) *delivery.AuthControllers{
	repo := repository.NewPostgresAuthRepo(db,rds)
	service := usecase.NewUserService(repo,rds)
	controllers := delivery.NewUserController(service)
	return controllers
}


