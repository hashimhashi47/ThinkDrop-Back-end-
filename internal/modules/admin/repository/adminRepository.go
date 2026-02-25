package repository

import (
	Domain "thinkdrop-backend/internal/Common"
	"thinkdrop-backend/internal/modules/admin/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) domain.AdminRepo {
	return &Repository{DB: db}
}

func (r *Repository) GetCounts(Data string) (int64, error) {
	var count int64
	err := r.DB.Table(Data).Count(&count).Error
	return count, err
}

func (r *Repository) FindTopUser(data interface{}) error {
	return r.DB.Order("points_available DESC").Limit(1).First(data).Error
}

func (r *Repository) FindbyArgs(data interface{}, query string, args interface{}) error {
	return r.DB.Where(query, args).First(data).Error
}

func (r *Repository) FindAll(model interface{}, limit, offset int, preload1, preload2 string) error {
	return r.DB.Limit(limit).Offset(offset).Preload(preload1).Preload(preload2).Order("created_at DESC").Find(model).Error
}

func (r *Repository) Find(model interface{}) error {
	return r.DB.Find(model).Error
}

func (r *Repository) FindWithoutPreload(model interface{}, limit, offset int) error {
	return r.DB.Limit(limit).Offset(offset).Order("created_at DESC").Find(model).Error
}

func (r *Repository) Count(model interface{}) (int64, error) {
	var total int64
	if err := r.DB.Model(model).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *Repository) UpdateColumn(model interface{}, query string, id interface{}, column string, value interface{},
) error {

	return r.DB.
		Model(model).
		Where(query, id).
		UpdateColumn(column, value).
		Error
}

func (r *Repository) FindReportedPosts(model interface{}, limit, offset int, preloads ...string) error {
	db := r.DB.Where("report_count > ?", 20).
		Order("report_count DESC").
		Limit(limit).
		Offset(offset)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	return db.Find(model).Error
}

func (r *Repository) FindWithPreload(model interface{}, limit, offset int, preload string) error {
	return r.DB.Limit(limit).Offset(offset).Preload(preload).Order("created_at DESC").Find(model).Error
}

func (r *Repository) Delete(model interface{}, query string, arg interface{}) error {
	return r.DB.Unscoped().Where(query, arg).Delete(model).Error
}

func (r *Repository) Insert(model interface{}) error {
	return r.DB.Create(model).Error
}

func (r *Repository) DeletePostWithRelations(postID uint) error {

	tx := r.DB.Begin()

	var post Domain.Post

	// Find post
	if err := tx.First(&post, postID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Clear many2many relation (join table)
	if err := tx.Model(&post).
		Association("SubInterests").
		Clear(); err != nil {
		tx.Rollback()
		return err
	}

	// Delete post permanently
	if err := tx.Unscoped().
		Delete(&post).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
