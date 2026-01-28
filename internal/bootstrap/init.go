package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery/authcontrollers"
	"thinkdrop-backend/internal/modules/auth/userAuth/delivery/otpcontrollers"
	"thinkdrop-backend/internal/modules/auth/userAuth/repository/databaseRepository"
	otprepository "thinkdrop-backend/internal/modules/auth/userAuth/repository/otpRepository"
	"thinkdrop-backend/internal/modules/auth/userAuth/usecase/authService"
	otpservice "thinkdrop-backend/internal/modules/auth/userAuth/usecase/otpService"
)

// ->Init the auth service to pass the connections from database to controllers
func InitAuth(db *gorm.DB, rds *redis.Client) *authcontrollers.UserController {
	repo := databaserepository.NewPostgresAuthRepo(db)
	service := authservice.NewUserService(repo, rds)
	controllers := authcontrollers.NewUserController(service)
	return controllers
}

// ->Init the otp service to pass the connection from redis to controllers
func InitOTP(rds *redis.Client) *otpcontrollers.OtpControllers {
	repo := otprepository.NewOTPrepository(rds)
	service := otpservice.NewOtpServices(repo)
	controllers := otpcontrollers.NewOtpServices(service)
	return controllers
}
