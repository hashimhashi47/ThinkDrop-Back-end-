package domain

import domain "thinkdrop-backend/internal/Common"

type PostRepo interface {
	FindAnything(model interface{}, Query, Any interface{}) error
	Insert(model interface{}) error
	AllowPost(userID uint) (bool, error)
	FindAnyWithpreload(model interface{}, Query, AnyData interface{}, Preload string) error
	FindByUser(model interface{}, Query string, Any interface{}) error
	FindFeedPosts(posts *[]domain.Post, subIDs []uint, limit int, offset int) error
}
