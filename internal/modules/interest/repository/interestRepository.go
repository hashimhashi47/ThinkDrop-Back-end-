package repository

import (
	domain "thinkdrop-backend/internal/Common"
	IntrestDomain "thinkdrop-backend/internal/modules/interest/domain"

	"gorm.io/gorm"
)

type InterestRepository struct {
	DB *gorm.DB
}

func NewPostgreIntrestRepo(db *gorm.DB) IntrestDomain.InterestRepo {
	return &InterestRepository{DB: db}
}

// -> get the entire intrest
func (r *InterestRepository) GetAll(model interface{}) error {
	return r.DB.Preload("SubInterests").Find(model).Error
}

// -> find by model,query with any like email,id etc 
func (r *InterestRepository) FindBy(model interface{}, Query string, Any interface{}) error {
	return r.DB.Where(Query, Any).Preload("SelectedSubs.MainInterest").First(model).Error
}

// -> save the model
func (r *InterestRepository) Save(model interface{}) error {
	return r.DB.Save(model).Error
}


// -> upadte the user intrests
func (r *InterestRepository) UpdateUserInterests(user interface{}, subIDs []uint) error {
    var subs []domain.SubInterest
    if err := r.DB.Where("id IN ?", subIDs).Find(&subs).Error; err != nil {
        return err
    }
    return r.DB.Model(user).Association("SelectedSubs").Replace(subs)
}
