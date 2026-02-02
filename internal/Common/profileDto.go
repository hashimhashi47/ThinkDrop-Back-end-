package domain

import "time"

type ProfilePage struct {
	AnonymousName  string
	bio            string
	FollowersCount int `gorm:"default:0"`
	FollowingCount int `gorm:"default:0"`
	WritingsCount  int `gorm:"default:0"`
}

type ProfileResponseDTO struct {
	AnonymousName string       `json:"anonymous_name"`
	WritingsCount int          `json:"writings_count"`
	Bio           string       `json:"bio"`
	Writings      []WritingDTO `json:"writings"`
}

type WritingDTO struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Intrests  string    `json:"intrest"`
	CreatedAt time.Time `json:"created_at"`
}
