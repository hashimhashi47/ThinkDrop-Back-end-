package usecase

import (
	"errors"
	"fmt"
	"log"
	"strings"
	domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/internal/utils"
)

// -> add account service logic

func (r *RewardService) AddUserAcoountService(UserID uint, Input domain.BankAccountInput) (domain.BankAccount, error) {
	var BankAccount domain.BankAccount
	var User domain.User

	if err := r.repo.Find(&BankAccount, "user_id = ?", UserID); err == nil {
		return domain.BankAccount{},
			errors.New("User already have a existing bank account")
	}

	if err := r.repo.Find(&User, "id = ?", UserID); err != nil {
		return domain.BankAccount{},
			errors.New("failed to find the user")
	}

	if Input.AccountNumber != Input.ReAccountNumber {
		return domain.BankAccount{},
			errors.New("account numbers do not match")
	}

	var bankPrefixes = map[string]string{
		"HDFC": "HDFC Bank",
		"SBIN": "State Bank of India",
		"ICIC": "ICICI Bank",
		"YESB": "Yes Bank",
	}

	BankAccount.BankName = "Unknown Bank"

	for prefix, name := range bankPrefixes {
		if strings.HasPrefix(strings.ToUpper(Input.IFSCCode), prefix) {
			BankAccount.BankName = name
			break
		}
	}

	razorContact, err := utils.RazorpayContact(User)
	if err != nil {
		return domain.BankAccount{}, err
	}

	BankAccount = domain.BankAccount{
		UserID:            UserID,
		AccountHolderName: Input.AccountHolderName,
		AccountNumber:     Input.AccountNumber,
		IFSCCode:          Input.IFSCCode,
		RazorpayContactID: razorContact.ID,
		IsVerified:        false,
	}

	razorFundAccount, err := utils.RazorpayFundAccount(BankAccount, User.FullName)

	if err != nil {
		return domain.BankAccount{}, err
	}

	BankAccount.RazorpayFundAccountID = razorFundAccount.ID

	if err := r.repo.Create(&BankAccount); err != nil {
		return domain.BankAccount{}, err
	}

	if err := r.repo.Update(&domain.Wallet{}, "user_id = ?", UserID, map[string]interface{}{
		"bank_account_id": BankAccount.ID}); err != nil {
		return domain.BankAccount{},
			errors.New("failed to link the account with wallet")
	}

	return BankAccount, nil
}

// -> withdraw fund logic service
func (r *RewardService) WithdrawPointsService(UserID uint, WithdrawPoints int64) (interface{}, error) {
	var Wallet domain.Wallet

	if err := r.repo.Find(&Wallet, "user_id = ?", UserID, "BankAccount"); err != nil {
		return nil, errors.New("failed to find the wallet")
	}

	// if !Wallet.BankAccount.IsVerified {
	// 	return nil, errors.New("bank account inactive")
	// }

	if WithdrawPoints > int64(Wallet.PointsAvailable) {
		return nil, errors.New("Your wallet have enough points")
	}

	if WithdrawPoints < 1000 {
		return nil, errors.New("only withdraw more than 1000 points")
	}

	var rate float64 = 0.01
	PointsToCash := float64(WithdrawPoints) * rate
	paise := PointsToCash * 100

	Data, err := utils.RazorpayPayout(int64(paise), Wallet.BankAccount.RazorpayFundAccountID)

	if err != nil {
		return nil, err
	}

	var Withdrawal domain.Withdrawal

	Withdrawal = domain.Withdrawal{
		UserID:           UserID,
		BankAccountID:    *Wallet.BankAccountID,
		PointsUsed:       int(WithdrawPoints),
		AmountINR:        int64(PointsToCash),
		Status:           Data.Status,
		RazorpayPayoutID: Data.ID,
		FailureReason:    Data.FailureReason,
	}

	if err := r.repo.Create(&Withdrawal); err != nil {
		return nil, errors.New("failed to create the withdrwal")
	}

	likesToDeduct := int(WithdrawPoints) / 2

	Wallet.PointsAvailable -= int(WithdrawPoints)
	Wallet.TotalLikes -= likesToDeduct

	if err := r.repo.Save(&Wallet); err != nil {
		return nil, err
	}

	AccountData, _ := r.AdminService.AddAccountStatusService()
	fmt.Println(AccountData)
	r.AdminService.Broadcast("accounts", "UPDATE_ACCOUNTS", AccountData)

	return Data, nil
}

func (r *RewardService) GetWithdrawalsService(UserID uint, limit, offset int) ([]domain.Withdrawal, error) {
	var withdrawals []domain.Withdrawal

	log.Println("limit", limit, "offset", offset)

	if err := r.repo.FindAllwithArg(&withdrawals, "user_id = ?", UserID, limit, offset); err != nil {
		return nil, errors.New("failed to find the withdrawals")
	}

	AccountData, _ := r.AdminService.AddAccountStatusService()
	WalletData, _ := r.AdminService.GetWalletsService(10, 0)
	r.AdminService.Broadcast("accounts", "UPDATE_ACCOUNTS", AccountData)
	r.AdminService.Broadcast("wallets", "UPDATE_WALLETS", WalletData)
	return withdrawals, nil
}

func (r *RewardService) GetWithdrawalsRefershService(UserID uint, limit, offset int) ([]domain.Withdrawal, error) {
	var withdrawals []domain.Withdrawal

	if err := r.repo.FindAllwithArg(&withdrawals, "user_id = ?", UserID, limit, offset); err != nil {
		return nil, errors.New("failed to find the withdrawals")
	}

	for i, w := range withdrawals {
		if w.Status == "processing" {
			payout, err := utils.GetRazorpayPayout(w.RazorpayPayoutID)
			if err != nil {
				continue
			}

			if payout.Status != w.Status {
				err := r.repo.Update(
					&domain.Withdrawal{},
					"razorpay_payout_id = ?",
					w.RazorpayPayoutID,
					map[string]interface{}{
						"status":         payout.Status,
						"failure_reason": payout.FailureReason,
						"utr":            payout.UTR,
					},
				)
				if err != nil {
					return nil, err
				}

				withdrawals[i].Status = payout.Status
				withdrawals[i].FailureReason = payout.FailureReason
				withdrawals[i].UTR = payout.UTR
			}
		}
	}

	AccountData, _ := r.AdminService.AddAccountStatusService()
	fmt.Println(AccountData)
	r.AdminService.Broadcast("accounts", "UPDATE_ACCOUNTS", AccountData)
	return withdrawals, nil
}
