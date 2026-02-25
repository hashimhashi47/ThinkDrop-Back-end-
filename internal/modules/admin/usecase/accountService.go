package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	DomainAdmin "thinkdrop-backend/internal/modules/admin/domain"
)

// -> logic for get the account status
func (a *AdminService) AddAccountStatusService() (interface{}, error) {
	var Wallets []domain.Wallet
	var Transaction []domain.Withdrawal

	var totalpoints int
	var totalpointsRedeemed int

	var response DomainAdmin.AccountResponse

	if err := a.repo.Find(&Wallets); err != nil {
		return nil, errors.New("failed to find the wallets of user")
	}

	if err := a.repo.Find(&Transaction); err != nil {
		return nil, errors.New("failed to find the wallets of user")
	}

	for _, v := range Wallets {
		totalpoints += v.PointsAvailable
	}

	for _, v := range Transaction {
		if v.Status == "processed" {
			totalpointsRedeemed += v.PointsUsed
		}
	}

	response = DomainAdmin.AccountResponse{
		TotalPointsAvailable: totalpoints,
		TotalRedeemed:        totalpointsRedeemed,
	}

	return response, nil

}

// -> logic for get all transaction logic
func (a *AdminService) GetWithdrawalsService(limit, offset int) (interface{}, int64, error) {
	var transactions []domain.Withdrawal

	//  Get total count
	total, err := a.repo.Count(&domain.Withdrawal{})
	if err != nil {
		return nil, 0, errors.New("failed to count transactions")
	}

	if err := a.repo.FindWithoutPreload(&transactions, limit, offset); err != nil {
		return nil, 0, errors.New("failed to get transactions")
	}

	return transactions, total, nil
}

// -> add main intrest
func (a *AdminService) AddMainIntrestService(Input DomainAdmin.MainCategory) error {
	MainIntrest := domain.MainInterest{
		Name: Input.CategoryName,
	}

	if err := a.repo.Insert(&MainIntrest); err != nil {
		return errors.Join(err, errors.New("failed to add main intrest"))
	}

	return nil
}

// -> upadte existing main intrest
func (a *AdminService) UpadteMainIntrestService(Input DomainAdmin.MainCategory, IntrestID uint) error {

	if err := a.repo.UpdateColumn(&domain.MainInterest{}, "id = ?", IntrestID, "Name",
		Input.CategoryName); err != nil {
		return err
	}

	return nil
}

// -> add sun intrest inside the sunintrest
func (a *AdminService) AddSubIntrestService(Input DomainAdmin.CreateSubInterestRequest) error {
	SubIntrest := domain.SubInterest{
		MainInterestID: Input.ParentID,
		Name:           Input.Name,
	}

	if err := a.repo.Insert(&SubIntrest); err != nil {
		return errors.Join(err, errors.New("failed to add sub intrest"))
	}

	return nil
}

// -> update the existing subintrest
func (a *AdminService) UpdateSubIntrestService(Input DomainAdmin.SubInterest, IntrestID uint) error {
	if err := a.repo.UpdateColumn(&domain.SubInterest{}, "id = ?", IntrestID, "Name",
		Input.SubInterestName); err != nil {
		return errors.Join(err, errors.New("failed to update subintrest"))
	}

	return nil
}

func (a *AdminService) DeleteIntrestService(MainID int) error {

	if MainID == 11 {
		return errors.New("failed to delete the others session")
	}

	if err := a.repo.Delete(&domain.MainInterest{}, "id = ?", MainID); err != nil {
		return errors.Join(err, errors.New("failed to delete the main intrest"))

	}

	return nil
}

func (a *AdminService) DeleteSubIntrestService(SubID int) error {

	if SubID == 37 {
		return errors.New("failed to delete the others session")
	}

	if err := a.repo.Delete(&domain.SubInterest{}, "id = ?", SubID); err != nil {
		return errors.Join(err, errors.New("failed to delete the Sub intrest"))

	}

	return nil
}
