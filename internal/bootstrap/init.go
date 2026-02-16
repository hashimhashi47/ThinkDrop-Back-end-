package bootstrap

import (
	AdminDelivery "thinkdrop-backend/internal/modules/admin/delivery"
	AdminRepository "thinkdrop-backend/internal/modules/admin/repository"
	AdminUsecase "thinkdrop-backend/internal/modules/admin/usecase"
	AdminHub "thinkdrop-backend/internal/modules/admin/websocket"

	AuthDelivery "thinkdrop-backend/internal/modules/auth/userAuth/delivery"
	AuthRepository "thinkdrop-backend/internal/modules/auth/userAuth/repository"
	AuthUsecase "thinkdrop-backend/internal/modules/auth/userAuth/usecase"

	ChatDelivery "thinkdrop-backend/internal/modules/chat/delivery"
	ChatRepository "thinkdrop-backend/internal/modules/chat/repository"
	ChatUsecase "thinkdrop-backend/internal/modules/chat/usecase"
	ChatHub "thinkdrop-backend/internal/modules/chat/websocket"

	ProfileDelivery "thinkdrop-backend/internal/modules/profile_page/delivery"
	ProfileRepository "thinkdrop-backend/internal/modules/profile_page/repository"
	ProfileUsecase "thinkdrop-backend/internal/modules/profile_page/usecase"

	RewardDelivery "thinkdrop-backend/internal/modules/reward/delivery"
	RewardRepository "thinkdrop-backend/internal/modules/reward/repository"
	RewardUsecase "thinkdrop-backend/internal/modules/reward/usecase"

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

func InitPost(db *gorm.DB, rds *redis.Client) *PostDelivery.PostControllers {
	repo := PostRepository.NewPostRepository(db, rds)
	service := PostUsecase.NewPostService(repo)
	controllers := PostDelivery.NewPostControllers(service)
	return controllers
}

func InitProfile(db *gorm.DB) *ProfileDelivery.ProfileController {
	repo := ProfileRepository.NewProfileRepository(db)
	service := ProfileUsecase.NewProfileService(repo)
	controllers := ProfileDelivery.NewProfileControllers(service)
	return controllers
}

func InitRewards(db *gorm.DB) *RewardDelivery.RewardController {
	repo := RewardRepository.NewRewardRepository(db)
	service := RewardUsecase.NewRewardService(repo)
	controllers := RewardDelivery.NewRewardController(service)
	return controllers
}

func InitChat(db *gorm.DB) *ChatDelivery.ChatHandler {
	repo := ChatRepository.NewChatRepository(db)
	service := ChatUsecase.NewChatService(repo)
	hub := ChatHub.NewHub()
	go hub.Run()
	return ChatDelivery.NewChatHandler(service, hub)
}

func InitAdmin(db *gorm.DB) *AdminDelivery.AdminController {
	repo := AdminRepository.NewAdminRepository(db)
	adminHub := AdminHub.NewHub()
	go adminHub.Run()
	Service := AdminUsecase.NewAdminService(repo, adminHub)
	return AdminDelivery.NewAdminController(Service, adminHub)
}
