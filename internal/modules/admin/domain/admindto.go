package domain

import "time"

type AccountResponse struct {
	TotalPointsAvailable int `json:"totalpoints"`
	TotalRedeemed        int `json:"totalRedeemed"`
}

type AdminUserDTO struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Verify    bool      `json:"verify"`
	IsBlocked bool      `json:"is_blocked"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminFlaggedPostResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	ReportCount int  `json:"report_count"`
	IsBlocked   bool `json:"is_blocked"`

	User    AdminPostUserDTO `json:"user"`
	Reports []AdminReportDTO `json:"reports"`
}

type AdminPostUserDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	AnonymousName string `json:"anonymous_name"`
	AvatarURL     string `json:"avatar_url"`
}

type AdminReportDTO struct {
	ID          uint      `json:"id"`
	ReportedBy  uint      `json:"reported_by"`
	Reason      string    `json:"reason"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type AdminProfile struct {
	ID        uint   `json:"id"`
	Name      string `json:"fullname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt time.Time
	ImageURL  string `json:"avatarurl"`
}
