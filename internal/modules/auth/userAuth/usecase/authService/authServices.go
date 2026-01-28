package authservice

import (
	"errors"
	"thinkdrop-backend/internal/config/redis"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/entity"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/repository/authRepository"
	"thinkdrop-backend/pkg/hashPass"
	"thinkdrop-backend/pkg/jwt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// → Auth business rules (services)

type UserServices struct {
	repo authrepository.AuthRespository
	rds  *goredis.Client
}

func NewUserService(r authrepository.AuthRespository, rd *goredis.Client) *UserServices {
	return &UserServices{repo: r, rds: rd}
}

// -> User Signup service bussiness logics
func (r *UserServices) UserSignupService(userDetails *entity.UserValidate) (user *entity.User, err error) {

	hashedPass, err := hashpass.GenerateHashedPassword(userDetails.Password)
	if err != nil {
		return nil, err
	}

	User := &entity.User{
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

func (r *UserServices) UserLoginService(UserLoginCredential *entity.Login) (user entity.User, AccessToken, RefereshTokenn string, err error) {
	var userDetails entity.User

	if err := r.repo.FindAnything(&userDetails, "email = ?", UserLoginCredential.Email); err != nil {
		return entity.User{}, "", "", errors.New("User not found")
	}

	if err := hashpass.CompareHashedPassword(UserLoginCredential.Password, userDetails.Password); err != nil {
		return entity.User{}, "", "", errors.New("Invalid password")
	}

	Accesstoken, err := jwt.AccessToken(userDetails.ID, user.Email, user.AnonymousName)

	if err != nil {
		return entity.User{}, "", "", errors.New("Failed to Create AccessToken")
	}

	RefereshToken, err := jwt.RefershToken(userDetails.ID, user.Email, user.AnonymousName)

	if err != nil {
		return entity.User{}, "", "", errors.New("Failed to Create RefershToken")
	}

	Key := "RefershToken:" + user.Email
	if err := r.rds.Set(redis.Ctx, Key, RefereshToken, 7*24*time.Hour).Err(); err != nil {
		return entity.User{}, "", "", errors.New("failed to store RefershToken")
	}

	return userDetails, Accesstoken, RefereshToken, nil
}
