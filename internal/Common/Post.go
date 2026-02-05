package domain

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Content string `json:"content" gorm:"type:text;not null"`

	// Foreign Keys
	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	SubInterestID uint        `json:"intrestid"`
	SubInterest   SubInterest `gorm:"foreignKey:SubInterestID"`

	// Counters (fast reads)
	LikeCount   int `gorm:"default:0"`
	ReportCount int `gorm:"default:0"`

	// Relations
	Comments []Comment
	Reports  []Report
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time

	Content string `gorm:"type:text;not null"`

	PostID uint
	Post   Post

	UserID uint
	User   User
}

type Report struct {
	ID uint `gorm:"primaryKey"`

	PostID uint
	UserID uint

	Reason string `gorm:"type:text"`

	CreatedAt time.Time
}
