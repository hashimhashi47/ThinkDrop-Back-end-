package domain

type ProfileRepo interface {
	Find(model interface{}, query string, args interface{}, preloads ...string) error
	Create(model interface{}) error
	CountFollow(followerID uint, followedID uint) (int64, error)
	Unfollow(userID uint, otherUserID uint) error
	FindAll(model interface{}) error 
	Save(model interface{}) error 
}
