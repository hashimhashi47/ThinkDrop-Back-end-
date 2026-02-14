package domain

import domain "thinkdrop-backend/internal/Common"

type ProfileRepo interface {
	Find(model interface{}, query string, args interface{}, preloads ...string) error
	GetUserPosts(userID uint, limit int, offset int) ([]domain.Post, error)
	CountFollow(followerID uint, followedID uint) (int64, error)
	Unfollow(userID uint, otherUserID uint) error
	FindAll(model interface{}) error
	Create(model interface{}) error 
	Save(model interface{}) error
}
