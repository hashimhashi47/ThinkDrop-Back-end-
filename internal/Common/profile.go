package domain

import "time"

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
