package domain

import "time"

type Message struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ConversationID uint      `gorm:"not null;index" json:"conversation_id"`

	SenderID uint `gorm:"not null;index" json:"sender_id"`

	Content string `gorm:"type:text;not null" json:"content"`

	IsSeen bool `gorm:"default:false" json:"is_seen"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type Conversation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	User1ID   uint      `gorm:"not null" json:"user1_id"`
	User2ID   uint      `gorm:"not null" json:"user2_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}