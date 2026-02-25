package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	AdminUsecase "thinkdrop-backend/internal/modules/admin/usecase"
	RewardDomain "thinkdrop-backend/internal/modules/reward/domain"
	"thinkdrop-backend/pkg/constants"
	walletid "thinkdrop-backend/pkg/walletID"
)

type RewardService struct {
	repo         RewardDomain.RewardRepo
	AdminService AdminUsecase.AdminService
}

func NewRewardService(r RewardDomain.RewardRepo, AS AdminUsecase.AdminService) *RewardService {
	return &RewardService{repo: r, AdminService: AS}
}

// -> wallet status chechking
func (r *RewardService) GetRewardStatusService(UserID uint) (bool, error) {
	var Wallet domain.Wallet

	if err := r.repo.Find(&Wallet, "user_id = ?", UserID, "BankAccount"); err != nil {
		return false, nil
	}

	return true, nil
}

// -> create wallet for that user
func (r *RewardService) CreateWalletService(userID uint) (domain.Wallet, error) {
	WalletID, err := walletid.GenerateWalletID("WALLET")

	if err != nil {
		return domain.Wallet{}, err
	}

	Wallet := domain.Wallet{
		WalletID:       WalletID,
		UserID:         userID,
		IsWalletActive: constants.WalletActive,
	}

	if err := r.repo.Create(&Wallet); err != nil {
		return domain.Wallet{}, errors.New("failed to create the wallet")
	}

	return Wallet, nil
}

// -> get the updated reward details
func (r *RewardService) GetRewardDetailsService(UserID uint) (domain.Wallet, error) {
	var User domain.User
	var wallet domain.Wallet

	if err := r.repo.Find(&wallet, "user_id = ?", UserID, "BankAccount"); err != nil {
		return domain.Wallet{}, errors.New("failed to find the wallet")
	}

	if err := r.repo.Find(&User, "id = ?", UserID, "Posts"); err != nil {
		return domain.Wallet{}, errors.New("failed to find the user")
	}

	var totalLike int

	for _, v := range User.Posts {
		totalLike += v.LikeCount
	}


	wallet.TotalLikes = totalLike

	return wallet, nil
}
