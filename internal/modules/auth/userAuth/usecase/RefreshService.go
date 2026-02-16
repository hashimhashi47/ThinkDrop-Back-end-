package usecase

import (
	"os"
	Redis "thinkdrop-backend/internal/config/redis"
	"thinkdrop-backend/pkg/constants"
	Pjwt "thinkdrop-backend/pkg/jwt"

	"github.com/golang-jwt/jwt/v5"
)

// ->
func (r *AuthService) RefereshTokenService(RefereshToken string) (string, error) {
	if RefereshToken == "" {
		return "", constants.ErrTokenNotFound
	}

	claim := &Pjwt.Claims{}
	tkn, err := jwt.ParseWithClaims(RefereshToken, claim, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !tkn.Valid {
		return "", constants.ErrTokenExpired
	}

	key := "RefershToken:" + claim.Email

	RdsToken, err := r.rds.Get(Redis.Ctx, key).Result()

	if err != nil || RdsToken != RefereshToken {
		return "", constants.ErrTokenMismatch
	}

	NewAccestoken, _ := Pjwt.AccessToken(claim.UserId, claim.Email, claim.AnonymousName, claim.Role)

	return NewAccestoken, nil
}
