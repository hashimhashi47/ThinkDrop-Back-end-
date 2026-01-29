package domain

import (
	"thinkdrop-backend/internal/modules/interest/domain"
	"time"

	"gorm.io/gorm"
)

// -> User enitire data struct
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FullName      string          `gorm:"not null"`
	AnonymousName string          `gorm:"uniqueIndex"`
	Verify        bool            `gorm:"false"`
	Email         string          `gorm:"uniqueIndex;not null"`
	Password      string          `gorm:"not null"`

	SelectedSubs []domain.SubInterest `gorm:"many2many:user_sub_interests;"`
}

// -> validation on user inputs
type UserValidate struct {
	FullName      string `json:"fullname" validate:"required"`
	AnonymousName string `json:"anonymousname" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Otp           string `json:"otp" validate:"required,len=6,numeric"`
	Password      string `json:"password" validate:"required,min=8,max=18"`
}

// -> login struct of user
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=18"`
}
