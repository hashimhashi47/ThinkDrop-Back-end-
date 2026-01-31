package repository

import (
	Redis "thinkdrop-backend/internal/config/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

// -> save the otp on redis
func (r *AuthRespository) SaveOTP(email, otp string) error {
	return r.redis.Set(Redis.Ctx, email, otp, 5*time.Minute).Err()
}

// -> Rate limiting the request it will avoid multiple requests
func (r *AuthRespository) RateLimitOTP(email string) (bool, error) {
	key := "OTP:RateLimit:" + email
	count, err := r.redis.Incr(Redis.Ctx, key).Result()
	if err != nil {
		return false, err
	}
	
	if count == 1 {
		r.redis.Expire(Redis.Ctx, key, 10*time.Minute)
	}

	if count > 3 {
		return false, nil
	}
	return true, nil
}

// -> Get the otp on redis
func (r *AuthRespository) GetOTP(email string) (OTP string, err error) {
	OTP, err = r.redis.Get(Redis.Ctx, email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return OTP, err
}

// -> delete the otp from redis
func (r *AuthRespository) DeleteOTP(email string) error {
	return r.redis.Del(Redis.Ctx, email).Err()
}
