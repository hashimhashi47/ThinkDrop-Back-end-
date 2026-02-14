package domain

// → Interface
type AuthRepo interface {
	Insert(model interface{}) error
	FindAnything(model interface{}, Query, Any interface{}) error
	SaveOTP(email, otp string) error
	RateLimitOTP(email string) (bool, error)
	DeleteOTP(email string) error
	GetOTP(email string) (OTP string, err error)
}
