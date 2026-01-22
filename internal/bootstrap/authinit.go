package bootstrap

import (
	"gorm.io/gorm"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	"thinkdrop-backend/internal/modules/auth/userAuth/repository"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase"
)

func InitAuth(db *gorm.DB) *delivery.UserController {
	repo := repository.NewPostgresAuthRepo(db)
	service := usecase.NewUserService(repo)
	controllers := delivery.NewUserController(service)
	return controllers
}
