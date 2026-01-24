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
	RandOTP := genrateotp.GenerateOTP()

	if err := utils.SentOTPEmail(email, RandOTP); err != nil {
		return "", err
	}

	if err := r.repo.SaveOTP(email, RandOTP); err != nil {
		return "", errors.Join(errors.New("failed to save the OTP"), err)
	}

	isOk, err := r.repo.RateLimitOTP(email)

	if err != nil {
		return "", err
	}

	if !isOk {
		return "", errors.New("Request limit exceeded, wait for 10 min")
	}

	return RandOTP, nil
}
