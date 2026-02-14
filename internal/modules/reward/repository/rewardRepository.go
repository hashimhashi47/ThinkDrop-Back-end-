package repository

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	RewardDomain "thinkdrop-backend/internal/modules/reward/domain"

	"gorm.io/gorm"
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

// -> find all with a specific id with limit and off set
func (r *RewardRepository) FindAllwithArg(model interface{}, query string, args interface{}, limit, offset int) error {
	return r.DB.
		Where(query, args).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(model).
		Error
}

// -> save the data˝
func (r *RewardRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}

// -> update anything
func (r *RewardRepository) Update(model interface{}, query string, arg interface{}, updates map[string]interface{},
) error {
	return r.DB.
		Model(model).
		Where(query, arg).
		Updates(updates).
		Error
}

func (r *RewardRepository) UpdateWallet(walletID uint, withdrawPoints int) error {
	if withdrawPoints <= 0 {
		return errors.New("invalid withdraw points")
	}

	likesToDeduct := withdrawPoints / 2

	tx := r.DB.
		Debug().
		Model(&domain.Wallet{}).
		Where("id = ? AND points_available >= ?", walletID, withdrawPoints).
		UpdateColumn("points_available", gorm.Expr("points_available - ?", withdrawPoints)).
		UpdateColumn("total_likes", gorm.Expr("total_likes - ?", likesToDeduct))

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insufficient points or wallet not found")
	}

	return nil
}
