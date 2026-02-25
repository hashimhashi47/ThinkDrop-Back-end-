package domain

import (
	"gorm.io/gorm"
	domain "thinkdrop-backend/internal/Common"
)

type PostRepo interface {
	FindAnything(model interface{}, Query, Any interface{}) error
	FindAll(model interface{}, query string, args interface{}) error
	Insert(model interface{}) error
	AllowPost(userID uint) (bool, error)
	FindAnyWithpreload(model interface{}, query string, args interface{}, preloads ...string) error
	FindByUser(model interface{}, Query string, Any interface{}) error
	FindFeedPosts(posts *[]domain.Post, subIDs []uint, limit int, offset int) error
	Save(model interface{}) error
	DeleteWhere(model interface{}, query string, args ...interface{}) (*gorm.DB, error)
	Create(model interface{}) error
	UpdateColumn(model interface{}, query string, id interface{}, column string, value interface{}) error
	ReportRateLimit(PostID string) (bool, error)
	UpdateUserWalletByPostID(postID uint, points int) error
	FindLikeByUserID(userID uint, postID uint) (bool, error)
}
