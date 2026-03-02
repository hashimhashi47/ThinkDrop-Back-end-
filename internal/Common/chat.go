package domain

import "time"

type Message struct {
	ID             uint      `json:"id"`
	ConversationID uint      `json:"conversation_id"`
	SenderID       uint      `json:"sender_id"`
	ReceiverID     uint      `json:"receiver_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

type Conversation struct {
	ID            uint
	User1ID       uint
	User1NAME     string
	User1ImageUrl string
	User2ID       uint
	User2NAME     string
	User2ImageUrl string
	CreatedAt     time.Time
}




type UserMini struct {
	AnonymousName string
	ImageURL      string
}
