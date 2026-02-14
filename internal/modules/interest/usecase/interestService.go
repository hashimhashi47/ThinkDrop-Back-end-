package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	IntrestDomain "thinkdrop-backend/internal/modules/interest/domain"
)

type InterestService struct {
	repo IntrestDomain.InterestRepo
}

func NewInterestService(r IntrestDomain.InterestRepo) *InterestService {
	return &InterestService{repo: r}
}

// -> get the entire intrest to show
func (r *InterestService) ShowIntrestsService() ([]domain.MainInterest, error) {
	var interests []domain.MainInterest

	if err := r.repo.GetAll(&interests); err != nil {
		return nil, errors.New("failed to find interests")
	}

	var response []domain.MainInterest

	for _, v := range interests {
		var subs []domain.SubInterest // 👈 RESET INSIDE LOOP

		for _, j := range v.SubInterests {
			subs = append(subs, domain.SubInterest{
				ID:   j.ID,
				Name: j.Name,
			})
		}

		response = append(response, domain.MainInterest{
			ID:           v.ID,
			Name:         v.Name,
			SubInterests: subs,
		})
	}

	return response, nil
}

// -> add users intrests on thier table
func (r *InterestService) UserAddInterstsService(UserID uint, Intrests domain.Req) (domain.User, error) {
	var User domain.User

	if err := r.repo.FindBy(&User, "id = ?", UserID); err != nil {
		return domain.User{}, err
	}

	if err := r.repo.UpdateUserInterests(&User, Intrests.SubInterestIDs); err != nil {
		return domain.User{}, err
	}

	if err := r.repo.Save(User); err != nil {
		return domain.User{}, err
	}

	if err := r.repo.FindBy(&User, "id = ?", UserID); err != nil {
		return domain.User{}, err
	}

	return User, nil
}
