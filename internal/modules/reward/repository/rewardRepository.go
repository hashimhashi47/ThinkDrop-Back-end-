package repository

import (
	"gorm.io/gorm"
	RewardDomain "thinkdrop-backend/internal/modules/reward/domain"
)

type RewardRepository struct {
	DB *gorm.DB
}

func NewRewardRepository(db *gorm.DB) RewardDomain.RewardRepo {
	return &RewardRepository{DB: db}
}

// -> we can make find somthing with multiple preload
func (r *RewardRepository) Find(model interface{}, query string, args interface{}, preloads ...string) error {
	db := r.DB.Where(query, args)
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db.First(model).Error
}

// -> create on database
func (r *RewardRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

// -> find all from a table
func (r *RewardRepository) FindAll(model interface{}) error {
	return r.DB.Find(model).Error
}

// -> save the data
func (r *RewardRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}

func (r *RewardRepository) Update(model interface{}, query string, arg interface{}, updates map[string]interface{},
) error {

	return r.DB.
		Model(model).
		Where(query, arg).
		Updates(updates).
		Error
}
