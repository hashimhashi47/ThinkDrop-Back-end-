package bootstrap

import (
	AuthDelivery "thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	AuthRepository "thinkdrop-backend/internal/modules/auth/userAuth/repository"
	AuthUsecase "thinkdrop-backend/internal/modules/auth/userAuth/usecase"

	InterestDelivery "thinkdrop-backend/internal/modules/interest/delivery"
	InterestRepository "thinkdrop-backend/internal/modules/interest/repository"
	InterestUsecase "thinkdrop-backend/internal/modules/interest/usecase"

	PostDelivery "thinkdrop-backend/internal/modules/post/delivery"
	PostRepository "thinkdrop-backend/internal/modules/post/repository"
	PostUsecase "thinkdrop-backend/internal/modules/post/usecase"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ->Init the auth service to pass the connections from database to controllers

func InitAuth(db *gorm.DB, rds *redis.Client) *AuthDelivery.AuthControllers {
	repo := AuthRepository.NewPostgresAuthRepo(db, rds)
	service := AuthUsecase.NewUserService(repo, rds)
	controllers := AuthDelivery.NewUserController(service)
	return controllers
}

func InitInterest(db *gorm.DB) *InterestDelivery.InterestControllers {
	repo := InterestRepository.NewPostgreIntrestRepo(db)
	service := InterestUsecase.NewInterestService(repo)
	controllers := InterestDelivery.NewInterestControllers(service)
	return controllers
}

func InitPost(db *gorm.DB,rds *redis.Client) *PostDelivery.PostControllers {
	repo := PostRepository.NewPostRepository(db,rds)
	service := PostUsecase.NewPostService(repo)
	controllers := PostDelivery.NewPostControllers(service)
	return controllers
}
