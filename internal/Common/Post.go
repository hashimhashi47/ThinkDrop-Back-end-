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

	SubInterests []SubInterest `gorm:"many2many:post_sub_interests;"`

	// Counters (fast reads)
	LikeCount    int `gorm:"default:0"`
	CommentCount int `gorm:"default:0"`
	ReportCount  int `gorm:"default:0"`

	// Relations
	Likes    []Like `gorm:"foreignKey:PostID"`
	Comments []Comment
	Reports  []Report
}

type Like struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;uniqueIndex:idx_user_post"`
	PostID uint `gorm:"not null;uniqueIndex:idx_user_post"`

	CreatedAt time.Time
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
