package repository

import (
	"gorm.io/gorm"
	ProfileDomain "thinkdrop-backend/internal/modules/profile_page/domain"
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
