package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	PostDomain "thinkdrop-backend/internal/modules/post/domain"
)

type PostService struct {
	repo PostDomain.PostRepo
}

func NewPostService(r PostDomain.PostRepo) *PostService {
	return &PostService{repo: r}
}

// -> user can upload their post it also have check limit request
func (r *PostService) AddPostService(Post domain.Post, UserID uint) (domain.Post, error) {
	var User domain.User
	var AddPost domain.Post

	if err := r.repo.FindAnything(&User, "id = ?", UserID); err != nil {
		return domain.Post{}, errors.New("failed to find user")
	}

	ok, err := r.repo.AllowPost(UserID)

	if err != nil {
		return domain.Post{}, errors.New("Too many request")
	}

	if !ok {
		return domain.Post{}, errors.New("try again after 30 seconds")
	}

	AddPost = domain.Post{
		Content:       Post.Content,
		UserID:        UserID,
		SubInterestID: Post.SubInterestID,
	}

	if err := r.repo.Insert(&AddPost); err != nil {
		return domain.Post{}, errors.New("Failed to post your Blog")
	}

	return AddPost, nil
}

// -> showAll users post can show their acccounts
func (r *PostService) ShowPostsServices(userID uint) ([]domain.PostResponse, error) {
	var posts []domain.Post

	if err := r.repo.FindAnyWithpreload(&posts, "user_id = ?", userID, "SubInterest"); err != nil {
		return nil, err
	}

	var response []domain.PostResponse

	for _, post := range posts {
		response = append(response, domain.PostResponse{
			ID:          post.ID,
			Content:     post.Content,
			SubInterest: post.SubInterest.Name,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		})
	}

	return response, nil
}

// -> user feed sesion it will customised by the intrest
func (r *PostService) UserFeedService(userID uint, limit, offset int) ([]domain.PostFeedResponse, error) {
	var User domain.User
	var AllPosts []domain.Post

	if err := r.repo.FindByUser(&User, "id = ?", userID); err != nil {
		return nil, err
	}

	subIDs := make([]uint, 0, len(User.SelectedSubs))
	for _, sub := range User.SelectedSubs {
		subIDs = append(subIDs, sub.ID)
	}

	if len(subIDs) == 0 {
		return []domain.PostFeedResponse{}, nil
	}

	if err := r.repo.FindFeedPosts(&AllPosts, subIDs, limit, offset); err != nil {
		return nil, err
	}

	feed := make([]domain.PostFeedResponse, 0, len(AllPosts))

	for _, post := range AllPosts {
		feed = append(feed, domain.PostFeedResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			LikeCount: post.LikeCount,
			Interest: domain.PostInterestDTO{
				PID:  post.SubInterestID,
				Name: post.SubInterest.Name,
			},
			User: domain.PostUserDTO{
				UID:           post.UserID,
				AnonymousName: post.User.AnonymousName,
			},
		})
	}
	return feed, nil
}
