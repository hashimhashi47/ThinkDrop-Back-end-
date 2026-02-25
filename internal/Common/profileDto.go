package domain

import "time"

type ProfilePage struct {
	AnonymousName  string
	Bio            string
	FollowersCount int `gorm:"default:0"`
	FollowingCount int `gorm:"default:0"`
	WritingsCount  int `gorm:"default:0"`
	Avatar         string
}

type ProfileResponseDTO struct {
	AnonymousName string      `json:"anonymous_name"`
	WritingsCount int          `json:"writings_count"`
	Bio           string       `json:"bio"`
	ImageURL      string       `json:"avatarurl"`
	Writings      []WritingDTO `json:"writings"`

	IsFollowing bool
}

type WritingDTO struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Intrests  []string  `json:"intrest"`
	ImageURL  string    `json:"avatarurl"`
	CreatedAt time.Time `json:"created_at"`
}

type UserProfileResponse struct {
	ID            uint
	FullName      string
	AnonymousName string
	Bio           string
	ProfileAvatar string
	Followers     []UserFollow
	Following     []UserFollow
}

func MapUserToProfile(u User) UserProfileResponse {
	return UserProfileResponse{
		ID:            u.ID,
		FullName:      u.FullName,
		AnonymousName: u.AnonymousName,
		Bio:           u.Bio,
		ProfileAvatar: u.ImageURL,
		Followers:     u.Followers,
		Following:     u.Following,
	}
}

type FollowUserResponse struct {
	UserID        uint    `json:"user_id"`
	AnonymousName string `json:"anonymous_name"`
	ImageURL      string  `json:"avatarurl"`

	IsFollowing bool `json:"is_following"` // I follow them
	IsFollower  bool `json:"is_follower"`  // They follow me
}
