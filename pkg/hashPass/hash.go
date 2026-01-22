package hashpass

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// -> Hashing Password
func GenerateHashedPassword(Password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Password hashing Failed")
	}
	return string(hash), nil
}

// -> Compare string and hashed password
func CompareHashedPassword(InputPassword, HashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(InputPassword)); err != nil {
		return errors.New("Password not matched")
	}
	return nil
}
