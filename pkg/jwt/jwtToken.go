package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId        uint   `json:"userid"`
	Email         string `json:"email"`
	AnonymousName string
	Role          string `json:"role"`
	jwt.RegisteredClaims
}

// -> AccessToken creation including user details and short period for expiring
func AccessToken(UserID uint, email, AnonumousName, role string) (string, error) {
	Key := []byte(os.Getenv("JWT_SECRET_KEY"))
	claim := Claims{
		Email:         email,
		UserId:        UserID,
		AnonymousName: AnonumousName,
		Role:          role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	encodeToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	AccessToken, err := encodeToken.SignedString(Key)
	return AccessToken, err
}

// ->  RefershToken creation including user details and Long period for expiring
func RefershToken(UserID uint, email, anonymousname, role string) (string, error) {
	Key := []byte(os.Getenv("JWT_SECRET_KEY"))

	claim := Claims{
		UserId:        UserID,
		Email:         email,
		AnonymousName: anonymousname,
		Role:          role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	encodeToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	RefershToken, err := encodeToken.SignedString(Key)
	return RefershToken, err
}
