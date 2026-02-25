package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
)

func (a *AdminService) GetWalletsService(limit, offset int) (interface{}, error) {
	var Wallet []domain.Wallet

	if err := a.repo.FindWithPreload(&Wallet, limit, offset, "BankAccount"); err != nil {
		return nil, errors.New("failed to find the wallet")
	}

	return Wallet, nil
}

func (a *AdminService) BlockWalletService(WalletID int) (interface{}, error) {
	var Wallet domain.Wallet

	if err := a.repo.UpdateColumn(&Wallet, "id = ?", WalletID, "IsWalletBlocked", true); err != nil {
		return nil, errors.New("failed to block the wallet")
	}
	return Wallet, nil
}

func (a *AdminService) UnBlockWalletService(WalletID int) (interface{}, error) {
	var Wallet domain.Wallet

	if err := a.repo.UpdateColumn(&Wallet, "id = ?", WalletID, "IsWalletBlocked", false); err != nil {
		return nil, errors.New("failed to block the wallet")
	}
	return Wallet, nil
}

func (a *AdminService) VerifyAccountService(BankID int) (interface{}, error) {
	var BankAccount domain.BankAccount

	if err := a.repo.UpdateColumn(&BankAccount, "id = ?", BankID, "IsVerified", true); err != nil {
		return nil, errors.New("failed to block the wallet")
	}
	return BankAccount, nil
}

