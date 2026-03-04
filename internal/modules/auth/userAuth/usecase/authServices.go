package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	Redis "thinkdrop-backend/internal/config/redis"
	AdminDomain "thinkdrop-backend/internal/modules/admin/domain"
	AdminUsecase "thinkdrop-backend/internal/modules/admin/usecase"
	AuthDomain "thinkdrop-backend/internal/modules/auth/userAuth/domain"
	"thinkdrop-backend/pkg/constants"
	hashpass "thinkdrop-backend/pkg/hashPass"
	"thinkdrop-backend/pkg/jwt"
	"time"

	"github.com/redis/go-redis/v9"
)

// → Auth business rules (services)

type AuthService struct {
	repo         AuthDomain.AuthRepo
	rds          *redis.Client
	AdminService AdminUsecase.AdminService
}

func NewUserService(r AuthDomain.AuthRepo, rd *redis.Client, As AdminUsecase.AdminService) *AuthService {
	return &AuthService{repo: r, rds: rd, AdminService: As}
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
		ImageURL:      "https://api.dicebear.com/7.x/lorelei/svg",
	}

	if err := r.repo.Insert(User); err != nil {
		return nil, errors.New("Signup failed")
	}

	Userdata, _, _ := r.AdminService.GetUsersDetailService(10, 0)
	r.AdminService.Broadcast("users", "UPDATE_USER", Userdata)
	return User, nil

}

func (r *AuthService) UserLoginService(UserLoginCredential *domain.Login) (user domain.User,
	AccessToken, RefereshTokenn string, Role interface{}, err error) {
	var userDetails domain.User

	if err := r.repo.FindAnything(&userDetails, "email = ?", UserLoginCredential.Email); err != nil {
		return domain.User{}, "", "", nil, errors.New("User not found")
	}

	if err := hashpass.CompareHashedPassword(UserLoginCredential.Password, userDetails.Password); err != nil {
		return domain.User{}, "", "", nil, errors.New("Invalid password")
	}

	Accesstoken, err := jwt.AccessToken(userDetails.ID, user.Email, user.AnonymousName, userDetails.Role)

	if err != nil {
		return domain.User{}, "", "", nil, errors.New("Failed to Create AccessToken")
	}

	RefereshToken, err := jwt.RefershToken(userDetails.ID, user.Email, user.AnonymousName, userDetails.Role)

	if err != nil {
		return domain.User{}, "", "", nil, errors.New("Failed to Create RefershToken")
	}

	Key := "RefershToken:" + user.Email
	if err := r.rds.Set(Redis.Ctx, Key, RefereshToken, 7*24*time.Hour).Err(); err != nil {
		return domain.User{}, "", "", nil, errors.New("failed to store RefershToken")
	}

	if userDetails.Role != constants.User {
		var Permisison AdminDomain.Role
		if errr := r.repo.FindAnythingProload(&Permisison, "name = ?", userDetails.Role); errr != nil {
			return domain.User{}, "", "", nil, errr
		}

		return userDetails, Accesstoken, RefereshToken, Permisison, nil
	}

	return userDetails, Accesstoken, RefereshToken, nil, nil
}

func (r *AuthService) LogoutService(userId uint) error {
	var User domain.User

	if err := r.repo.FindAnything(&User, "id = ?", userId); err != nil {
		return errors.New("failed to find teh user")
	}

	key := "RefershToken:" + User.Email

	if err := r.rds.Del(Redis.Ctx, key).Err(); err != nil {
		return errors.New("failed to delete refresh token")
	}

	return nil
}
