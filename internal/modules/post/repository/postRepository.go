package repository

import (
	"fmt"
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

func (r *PostRepository) FindAnything(model interface{}, Query, Any interface{}) error {
	return r.DB.Where(Query, Any).First(model).Error
}

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

func (r *PostRepository) FindAnyWithpreload(model interface{}, Query, AnyData interface{}, Preload string) error {
	return r.DB.Preload(Preload).Where(Query, AnyData).Find(model).Error
}

func (r *PostRepository) FindByUser(model interface{}, Query string, Any interface{}) error {
	return r.DB.Where(Query, Any).Preload("SelectedSubs.MainInterest").First(model).Error
}

func (r *PostRepository) FindAll(model interface{}) error {
	return r.DB.Find(model).Error
}


