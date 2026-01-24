package otpservice

import (
	"errors"
	otprepo "thinkdrop-backend/internal/modules/auth/userAuth/domain/repository/otpRepo"
	"thinkdrop-backend/internal/utils"
	genrateotp "thinkdrop-backend/pkg/genrateOTP"
)

type OtpService struct {
	repo otprepo.OTPrepo
}

func NewOtpServices(r otprepo.OTPrepo) *OtpService {
	return &OtpService{repo: r}
}

// -> OTP generate busniess logics
func (r *OtpService) SentOtpService(email string) (OTP string, err error) {

	isOk, err := r.repo.RateLimitOTP(email)

	if err != nil {
		return "", err
	}

	if !isOk {
		return "", errors.New("Request limit exceeded, wait for 10 min")
	}

	RandOTP := genrateotp.GenerateOTP()

	if err := utils.SentOTPEmail(email, RandOTP); err != nil {
		return "", err
	}

	if err := r.repo.SaveOTP(email, RandOTP); err != nil {
		return "", errors.Join(errors.New("failed to save the OTP"), err)
	}

	return RandOTP, nil
}

// ->Verify the email logic
func (r *OtpService) OTPverifyService(email, OTP string) error {
	storedOtp, err := r.repo.GetOTP(email)

	if err != nil {
		return errors.Join(errors.New("OTP mismatched"), err)
	}

	if storedOtp != OTP {
		return errors.New("invalid otp")
	}

	r.repo.DeleteOTP(email)
	return nil
}
