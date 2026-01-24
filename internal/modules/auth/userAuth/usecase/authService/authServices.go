package authservice

import (
	"errors"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/entity"
	"thinkdrop-backend/internal/modules/auth/userAuth/domain/repository/authRepository"
	"thinkdrop-backend/pkg/hashPass"
)

// → Auth business rules (services)

type UserServices struct {
	repo authrepository.AuthRespository
}

func NewUserService(r authrepository.AuthRespository) *UserServices {
	return &UserServices{repo: r}
}

// -> User login service bussiness logics
func (r *UserServices) UserLoginService(userDetails *entity.UserValidate) (user *entity.User, err error) {

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
