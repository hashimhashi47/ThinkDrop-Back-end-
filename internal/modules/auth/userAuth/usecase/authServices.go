package usecase

import (
	"errors"
	Redis "thinkdrop-backend/internal/config/redis"
	AuthDomain"thinkdrop-backend/internal/modules/auth/userAuth/domain"
	domain "thinkdrop-backend/internal/Common"
	hashpass "thinkdrop-backend/pkg/hashPass"
	"thinkdrop-backend/pkg/jwt"
	"time"

	"github.com/redis/go-redis/v9"
)

// → Auth business rules (services)

type AuthService struct {
	repo AuthDomain.AuthRepo
	rds  *redis.Client
}

func NewUserService(r AuthDomain.AuthRepo, rd *redis.Client) *AuthService {
	return &AuthService{repo: r, rds: rd}
}

// -> User Signup service bussiness logics
func (r *AuthService) UserSignupService(userDetails *domain.UserValidate) (user *domain.User, err error) {

	hashedPass, err := hashpass.GenerateHashedPassword(userDetails.Password)
	if err != nil {
		return nil, err
	}

	User := &domain.User{
		FullName:      userDetails.FullName,
		AnonymousName: userDetails.AnonymousName,
		Email:         userDetails.Email,
		Password:      hashedPass,
	}

	if err := r.repo.Insert(User); err != nil {
		return nil, errors.New("Signup failed")
	}

	return User, nil

}

func (r *AuthService) UserLoginService(UserLoginCredential *domain.Login) (user domain.User, AccessToken, RefereshTokenn string, err error) {
	var userDetails domain.User

	if err := r.repo.FindAnything(&userDetails, "email = ?", UserLoginCredential.Email); err != nil {
		return domain.User{}, "", "", errors.New("User not found")
	}

	if err := hashpass.CompareHashedPassword(UserLoginCredential.Password, userDetails.Password); err != nil {
		return domain.User{}, "", "", errors.New("Invalid password")
	}

	Accesstoken, err := jwt.AccessToken(userDetails.ID, user.Email, user.AnonymousName)

	if err != nil {
		return domain.User{}, "", "", errors.New("Failed to Create AccessToken")
	}

	RefereshToken, err := jwt.RefershToken(userDetails.ID, user.Email, user.AnonymousName)

	if err != nil {
		return domain.User{}, "", "", errors.New("Failed to Create RefershToken")
	}

	Key := "RefershToken:" + user.Email
	if err := r.rds.Set(Redis.Ctx, Key, RefereshToken, 7*24*time.Hour).Err(); err != nil {
		return domain.User{}, "", "", errors.New("failed to store RefershToken")
	}

	return userDetails, Accesstoken, RefereshToken, nil
}
