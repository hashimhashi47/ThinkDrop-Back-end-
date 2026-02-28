package domain

import (
	"thinkdrop-backend/pkg/constants"
	"time"

	"gorm.io/gorm"
)

// -> Follow of the user
type UserFollow struct {
	ID uint `gorm:"primaryKey"`

	FollowerID uint
	Follower   User `gorm:"foreignKey:FollowerID"`

	FollowedID uint
	Followed   User `gorm:"foreignKey:FollowedID"`

	CreatedAt time.Time
}

// -> avatars
type Avatar struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ImageURL string `gorm:"not null" json:"image_url"`
}

type EditProfile struct {
	AnonymousName string `json:"anonymous_name" validate:"omitempty,min=3,max=30"`
	ImageURL      string `json:"image_url" validate:"omitempty,url"`
	Bio           string `json:"bio" validate:"omitempty,max=160"`
}

type ReportComplaints struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Type        constants.ReportType `gorm:"type:varchar(20);not null"`
	Description string               `gorm:"type:text;not null"`
	Status      string               `gorm:"type:varchar(20);default:'pending'"`

	// Foreign Key
	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}
