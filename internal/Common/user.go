package domain

import (
	"time"

	"gorm.io/gorm"
)

// -> User enitire data struct
type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FullName      string `gorm:"not null"`
	AnonymousName string `gorm:"uniqueIndex"`
	Verify        bool   `gorm:"false"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"`
	
	ImageURL     string        `json:"image_url" validate:"omitempty,url"`
	Bio          string        `json:"bio" validate:"omitempty,max=160"`
	Following    []UserFollow  `gorm:"foreignKey:FollowerID"`
	Followers    []UserFollow  `gorm:"foreignKey:FollowedID"`
	Posts        []Post        `gorm:"foreignKey:UserID"`
	
	SelectedSubs []SubInterest `gorm:"many2many:user_sub_interests;"`
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
