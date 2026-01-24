package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FullName      string `gorm:"not null"`
	AnonymousName string `gorm:"uniqueIndex"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"`
}

type UserValidate struct {
	FullName      string `json:"fullname" validate:"required"`
	AnonymousName string `json:"anonymousname" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Otp           string `json:"otp" validate:"required,len=6,numeric"`
	Password      string `json:"password" validate:"required,min=8,max=18"`
}
