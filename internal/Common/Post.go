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
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	SubInterests []SubInterest `gorm:"many2many:post_sub_interests;"`

	// Counters (fast reads)
	LikeCount    int `gorm:"default:0"`
	CommentCount int `gorm:"default:0"`
	ReportCount  int `gorm:"default:0"`

	// Relations
	Likes    []Like    `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE;"`
	Reports  []Report  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`

	Blocked bool `json:"blocked" gorm:"default:false"`
}

type Like struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;uniqueIndex:idx_user_post"`
	PostID uint `gorm:"not null;uniqueIndex:idx_user_post"`

	Post Post `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time

	Content string `gorm:"type:text;not null"`

	PostID uint
	Post   Post `gorm:"constraint:OnDelete:CASCADE;"`

	UserID uint
	User   User
}

type Report struct {
	ID uint `gorm:"primaryKey"`

	PostID uint
	UserID uint

	Post Post `gorm:"constraint:OnDelete:CASCADE;"`

	Reason      string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	ActionTaken bool   `gorm:"default:false"`

	CreatedAt time.Time
}

type ReportPostRequest struct {
	PostID      uint   `json:"postId"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}
