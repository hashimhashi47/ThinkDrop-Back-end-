package usecase

import (
	"errors"
	"thinkdrop-backend/internal/modules/interest/domain"
	IntrestDomain "thinkdrop-backend/internal/modules/interest/domain"
)

type InterestService struct {
	repo IntrestDomain.InterestRepo
}

func NewInterestService(r IntrestDomain.InterestRepo) *InterestService {
	return &InterestService{repo: r}
}

func (r *InterestService) ShowIntrestsService() ([]domain.MainInterestResponse, error) {
	var interests []domain.MainInterest

	if err := r.repo.GetAll(&interests); err != nil {
		return nil, errors.New("failed to find intrests")
	}
	var response []domain.MainInterestResponse
	var subs []domain.SubInterestResponse

	for _, v := range interests {

		for _, j := range v.SubInterests {
			subs = append(subs, domain.SubInterestResponse{
				ID:   j.ID,
				Name: j.Name,
			})
		}

		response = append(response, IntrestDomain.MainInterestResponse{
			ID:           v.ID,
			Name:         v.Name,
			SubInterests: subs,
		})

	}

	return response, nil
}
