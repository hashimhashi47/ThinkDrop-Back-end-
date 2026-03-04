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

func (r *Repository) FindAllWithOnePreload(model interface{}, limit, offset int, preload string) error {
	return r.DB.Limit(limit).Offset(offset).Preload(preload).Order("created_at DESC").Find(model).Error
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

	if err := tx.First(&post, postID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&post).
		Association("SubInterests").
		Clear(); err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Unscoped().
		Delete(&post).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) FindPreload(model interface{}, preload string) error {
	return r.DB.Preload(preload).Order("created_at DESC").Find(model).Error
}


func (r *Repository) ReplaceRolePermissions(roleID uint, permissionIDs []uint) error {
	var role domain.Role

	if err := r.DB.First(&role, roleID).Error; err != nil {
		return err
	}

	var permissions []domain.Permission

	if len(permissionIDs) > 0 {
		if err := r.DB.
			Where("id IN ?", permissionIDs).
			Find(&permissions).Error; err != nil {
			return err
		}
	}

	if err := r.DB.
		Model(&role).
		Association("Permissions").
		Replace(&permissions); err != nil {
		return err
	}

	return nil
}


