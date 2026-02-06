package domain

import "time"

// -> feed DTO
type PostFeedResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	LikeCount int `json:"like_count"`

	Interest PostInterestDTO `json:"interest"`
	User     PostUserDTO     `json:"user"`
}

type PostUserDTO struct {
	UID           uint   `json:"id"`
	AnonymousName string `json:"anonymous_name"`
}

type PostInterestDTO struct {
	PID  uint   `json:"id"`
	Name string `json:"name"`
}

// -> Profile show post DTO
type PostResponse struct {
	ID           uint      `json:"id"`
	Content      string    `json:"content"`
	SubInterest  string    `json:"sub_interest"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LikeCount    int       `json:"likecount"`
	CommentCount int       `json:"commentcount"`
}
