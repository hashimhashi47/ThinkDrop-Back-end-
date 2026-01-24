package otprepo

type OTPrepo interface {
	SaveOTP(email, otp string) error
	RateLimitOTP(email string) (bool, error)
	DeleteOTP(email string) error
	GetOTP(email string) (OTP string, err error)
}
