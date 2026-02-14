package repository

import (
	"fmt"
	Commom "thinkdrop-backend/internal/Common"
	RedisP "thinkdrop-backend/internal/config/redis"
	"thinkdrop-backend/internal/modules/post/domain"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB  *gorm.DB
	RDS *redis.Client
}

func NewPostRepository(db *gorm.DB, rds *redis.Client) domain.PostRepo {
	return &PostRepository{DB: db, RDS: rds}
}

// -> find anything with query and what with need to chack
func (r *PostRepository) FindAnything(model interface{}, Query, Any interface{}) error {
	return r.DB.Where(Query, Any).First(model).Error
}

// -> insert data into table
func (r *PostRepository) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}

// -> it will block multiple requests at a time for a post
func (r *PostRepository) AllowPost(userID uint) (bool, error) {

	key := fmt.Sprintf("post_rate_limit:user:%d", userID)

	count, err := r.RDS.Incr(RedisP.Ctx, key).Result()

	if err != nil {
		return false, err
	}

	if count == 1 {
		r.RDS.Expire(RedisP.Ctx, key, 40*time.Second)
	}

	if count > 1 {
		return false, nil
	}

	return true, nil
}

// -> find anything include preload
func (r *PostRepository) FindAnyWithpreload(model interface{}, query string, args interface{}, preloads ...string) error {
	db := r.DB.Where(query, args)
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db.First(model).Error
}

func (r *PostRepository) FindAll(model interface{}, query string, args interface{}) error {
	return r.DB.Where(query, args).Find(model).Error
}

// -> finc user with preload
func (r *PostRepository) FindByUser(model interface{}, Query string, Any interface{}) error {
	return r.DB.Where(Query, Any).Preload("SelectedSubs.MainInterest").First(model).Error
}

// -> specifically for find the post releated to the user
func (r *PostRepository) FindFeedPosts(posts *[]Commom.Post, subIDs []uint, limit, offset int) error {
	return r.DB.
		Joins("JOIN post_sub_interests psi ON psi.post_id = posts.id").
		Where("psi.sub_interest_id IN ?", subIDs).
		Preload("User").
		Preload("SubInterests").
		Order("posts.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(posts).Error
}


func (r *PostRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}

func (r *PostRepository) DeleteWhere(model interface{}, query string, args ...interface{},
) (*gorm.DB, error) {
	result := r.DB.Where(query, args...).Delete(model)
	return result, result.Error
}

// -> create on database
func (r *PostRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

func (r *PostRepository) UpdateColumn(model interface{}, query string, id interface{}, column string, value interface{},
) error {

	return r.DB.
		Model(model).
		Where(query, id).
		UpdateColumn(column, value).
		Error
}
