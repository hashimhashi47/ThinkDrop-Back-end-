package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	PostDomain "thinkdrop-backend/internal/modules/post/domain"

	"gorm.io/gorm"
)

type PostService struct {
	repo PostDomain.PostRepo
}

func NewPostService(r PostDomain.PostRepo) *PostService {
	return &PostService{repo: r}
}

// -> user can upload their post it also have check limit request
func (r *PostService) AddPostService(Post domain.CreatePostRequest, UserID uint) (domain.Post, error) {
	var User domain.User
	var AddPost domain.Post
	var subInterests []domain.SubInterest

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

	if err := r.repo.FindAll(&subInterests, "id IN ?", Post.InterestIDs); err != nil {
		return domain.Post{}, errors.New("failed to fetch interests")
	}

	if len(subInterests) != len(Post.InterestIDs) {
		return domain.Post{}, errors.New("invalid interest ids")
	}

	AddPost = domain.Post{
		Content:      Post.Content,
		UserID:       UserID,
		SubInterests: subInterests,
	}

	if err := r.repo.Insert(&AddPost); err != nil {
		return domain.Post{}, errors.New("Failed to post your Blog")
	}

	return AddPost, nil
}

// -> show All users post can show their acccounts
func (r *PostService) ShowPostsServices(userID uint) ([]domain.PostResponse, error) {
	var posts []domain.Post

	if err := r.repo.FindAnyWithpreload(&posts, "user_id = ?", userID, "SubInterests"); err != nil {
		return nil, err
	}

	var response []domain.PostResponse

	for _, post := range posts {
		var interests []string

		for _, v := range post.SubInterests {
			interests = append(interests, v.Name)
		}
		response = append(response, domain.PostResponse{
			ID:           post.ID,
			Content:      post.Content,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			SubInterest:  interests,
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

		interests := make([]domain.PostInterestDTO, 0)

		for _, si := range post.SubInterests {
			interests = append(interests, domain.PostInterestDTO{
				PID:  si.ID,
				Name: si.Name,
			})
		}

		feed = append(feed, domain.PostFeedResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			LikeCount: post.LikeCount,
			Interests: interests,
			User: domain.PostUserDTO{
				UID:           post.UserID,
				AnonymousName: post.User.AnonymousName,
				ImageURL:      post.User.ImageURL,
			},
		})
	}
	return feed, nil
}

// -> add like on post logic
func (r *PostService) LikePostService(UserID uint, PostID int) (bool, error) {
	Like := domain.Like{
		UserID: UserID,
		PostID: uint(PostID),
	}

	if err := r.repo.Create(&Like); err != nil {
		return false, errors.New("already liked the post")
	}

	if err := r.repo.UpdateColumn(&domain.Post{}, "id = ?", PostID, "like_count",
		gorm.Expr("like_count + 1")); err != nil {
		return false, err
	}
	return true, nil
}

// -> unlike post
func (r *PostService) UnLikePostService(UserID uint, PostID int) (bool, error) {
	result, err := r.repo.DeleteWhere(&domain.Like{}, "user_id = ? AND post_id = ?", UserID, PostID)
	if err != nil {
		return false, err
	}

	if result.RowsAffected == 0 {
		return false, errors.New("post not liked")
	}

	if err := r.repo.UpdateColumn(&domain.Post{}, "id = ?", PostID, "like_count",
		gorm.Expr("like_count - 1")); err != nil {
		return false, err
	}

	return true, nil
}
