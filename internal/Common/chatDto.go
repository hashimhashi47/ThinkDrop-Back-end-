package domain

import "time"

type ChatSidebarDTO struct {
	ConversationID uint      `json:"conversation_id"`
	UserID         uint      `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserAvatar     string    `json:"user_avatar"`
	LastMessage    string    `json:"last_message"`
	LastMessageAt  time.Time `json:"last_message_at"`
	UnreadCount    int       `json:"unread_count"`
	IsSeen         bool      `json:"is_seen"`
}


type OtherUserDTO struct {
	ConversationID uint      `json:"conversation_id"`
	UserID         uint      `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserImageUrl   string    `json:"user_image_url"`
	CreatedAt      time.Time `json:"created_at"`
}