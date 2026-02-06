package usecase

import (
	"errors"
	"strings"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/internal/utils"
)

// -> add account service logic

func (r *RewardService) AddUserAcoountService(UserID uint, Input domain.BankAccountInput) (domain.BankAccount, error) {
	var BankAccount domain.BankAccount
	var User domain.User

	if err := r.repo.Find(&BankAccount, "user_id = ?", UserID); err != nil {
		return domain.BankAccount{},
			errors.New("User already have a existing bank account")
	}

	if err := r.repo.Find(&User, "id = ?", UserID); err != nil {
		return domain.BankAccount{},
			errors.New("failed to find the user")
	}

	if Input.AccountNumber != Input.ReAccountNumber {
		return domain.BankAccount{},
			errors.New("Mismatched account number")
	}

	var bankPrefixes = map[string]string{
		"HDFC": "HDFC Bank",
		"SBIN": "State Bank of India",
		"ICIC": "ICICI Bank",
		"YESB": "Yes Bank",
	}

	razorData, err := utils.RazorpayContact(User)
	if err != nil {
		return domain.BankAccount{}, err
	}

	BankAccount = domain.BankAccount{
		UserID:            UserID,
		AccountHolderName: Input.AccountHolderName,
		IFSCCode:          Input.IFSCCode,
		RazorpayContactID: razorData.ID,
	}

	// Determine Bank Name
	for prefix, name := range bankPrefixes {
		if strings.Contains(strings.ToUpper(BankAccount.IFSCCode), prefix) {
			BankAccount.BankName = name
			break
		}
	}

	return BankAccount, nil
}
