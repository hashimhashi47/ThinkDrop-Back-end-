package repository

import (
	domain "thinkdrop-backend/internal/Common"
	ProfileDomain "thinkdrop-backend/internal/modules/profile_page/domain"

	"gorm.io/gorm"
)

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileDomain.ProfileRepo {
	return &ProfileRepository{DB: db}
}

// -> we can make find somthing with multiple preload
func (r *ProfileRepository) Find(model interface{}, query string, args interface{}, preloads ...string) error {
	db := r.DB.Where(query, args)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db.First(model).Error
}


// -> get users all post
func (r *ProfileRepository) GetUserPosts(userID uint,limit int,offset int,) ([]domain.Post, error) {
	var posts []domain.Post
	err := r.DB.
		Where("user_id = ?", userID).
		Preload("SubInterests").
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error

	return posts, err
}


// -> create on database
func (r *ProfileRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}


// -> take the count
func (r *ProfileRepository) CountFollow(followerID uint, followedID uint) (int64, error) {
	var count int64
	err := r.DB.Model(&domain.UserFollow{}).Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Count(&count).Error
	return count, err
}


// -> unnfollow user
func (r *ProfileRepository) Unfollow(userID uint, otherUserID uint) error {
	return r.DB.Where("follower_id = ? AND followed_id = ?", userID, otherUserID).
		Delete(&domain.UserFollow{}).Error
}


// -> find all from a table
func (r *ProfileRepository) FindAll(model interface{}) error {
	return r.DB.Find(model).Error
}


// -> save the data
func (r *ProfileRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}
