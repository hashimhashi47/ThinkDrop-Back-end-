package repository

import (
	IntrestDomain "thinkdrop-backend/internal/modules/interest/domain"
	"gorm.io/gorm"
)

type InterestRepository struct {
	DB *gorm.DB
}

func NewPostgreIntrestRepo(db *gorm.DB) IntrestDomain.InterestRepo {
	return &InterestRepository{DB: db}
}

func (r *InterestRepository) GetAll(model interface{}) error {
	return r.DB.Preload("SubInterests").Find(model).Error
}
