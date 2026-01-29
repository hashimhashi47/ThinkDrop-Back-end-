package usecase

import (
	"errors"
	AuthDomain "thinkdrop-backend/internal/modules/auth/userAuth/domain"
	IntrestDomain "thinkdrop-backend/internal/modules/interest/domain"
)

type InterestService struct {
	repo IntrestDomain.InterestRepo
}

func NewInterestService(r IntrestDomain.InterestRepo) *InterestService {
	return &InterestService{repo: r}
}

// -> get the entire intrest to show
func (r *InterestService) ShowIntrestsService() ([]IntrestDomain.MainInterestResponse, error) {
	var interests []IntrestDomain.MainInterest

	if err := r.repo.GetAll(&interests); err != nil {
		return nil, errors.New("failed to find intrests")
	}
	var response []IntrestDomain.MainInterestResponse
	var subs []IntrestDomain.SubInterestResponse

	for _, v := range interests {

		for _, j := range v.SubInterests {
			subs = append(subs, IntrestDomain.SubInterestResponse{
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

func (r *InterestService) UserAddInterstsService(UserID uint, Intrests IntrestDomain.Req) (AuthDomain.User, error) {
	var User AuthDomain.User

	if err := r.repo.FindBy(&User, "id = ?", UserID); err != nil {
		return AuthDomain.User{}, err
	}

	if err := r.repo.UpdateUserInterests(&User, Intrests.SubInterestIDs); err != nil {
		return AuthDomain.User{}, err
	}
	
	if err := r.repo.Save(User); err != nil {
		return AuthDomain.User{}, err
	}

	if err := r.repo.FindBy(&User, "id = ?", UserID); err != nil {
		return AuthDomain.User{}, err
	}
	return User, nil
}
