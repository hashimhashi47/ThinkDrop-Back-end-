package domain

import "time"

// -> Follow of the user
type UserFollow struct {
	ID            uint `gorm:"primaryKey"`
	AnonymousName string
	FollowerID    uint `gorm:"not null"` // who follows
	FollowedID    uint `gorm:"not null"` // who is being followed
	CreatedAt     time.Time
}
