package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	ProfileDomain "thinkdrop-backend/internal/modules/profile_page/domain"
)

type ProfileService struct {
	repo ProfileDomain.ProfileRepo
}

func NewProfileService(r ProfileDomain.ProfileRepo) *ProfileService {
	return &ProfileService{repo: r}
}

// -> show the own user details
func (r *ProfileService) ShowProfileService(UserID uint) (domain.User, error) {
	return domain.User{}, nil
}


// -> swow other users profile details 
func (r *ProfileService) ShowOtherUserProfileService(id int) (domain.ProfileResponseDTO, error) {
	var UserDetails domain.User

	if err := r.repo.Find(&UserDetails, "id = ?", id, "Posts.SubInterest"); err != nil {
		return domain.ProfileResponseDTO{}, errors.New("failed to find the user")
	}

	writings := make([]domain.WritingDTO, 0, len(UserDetails.Posts))
	for _, v := range UserDetails.Posts {
		writings = append(writings, domain.WritingDTO{
			ID:        v.ID,
			Content:   v.Content,
			Intrests:  v.SubInterest.Name,
			CreatedAt: v.CreatedAt,
		})
	}

	return domain.ProfileResponseDTO{
		AnonymousName: UserDetails.AnonymousName,
		WritingsCount: len(UserDetails.Posts),
		Bio:           UserDetails.Bio,
		Writings:      writings,
	}, nil
}
