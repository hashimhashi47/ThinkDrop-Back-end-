package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	AdminDomain "thinkdrop-backend/internal/modules/admin/domain"
)

func (a *AdminService) BlockPostService(PostID uint) error {
	if err := a.repo.UpdateColumn(domain.Post{}, "id = ?", PostID, "Blocked", true); err != nil {
		return errors.Join(errors.New("failed to update the block"), err)
	}

	return nil
}

func (a *AdminService) UnBlockPostService(PostID uint) error {
	if err := a.repo.UpdateColumn(domain.Post{}, "id = ?", PostID, "Blocked", false); err != nil {
		return errors.Join(errors.New("failed to update the block"), err)
	}
	
	return nil
}

func (a *AdminService) EditInrtestService(PostID uint, Response AdminDomain.UpdatePostInterestRequest) error {

	if err := a.repo.UpdateColumn(domain.Post{}, "id = ?", PostID, "SubInterests", Response.InterestIDs); err != nil {
		return errors.Join(errors.New("failed to update the block"), err)
	}

	return nil
}
