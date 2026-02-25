package domain

import domain "thinkdrop-backend/internal/Common"

type ChatRepository interface {
	FindConversation(user1, user2 uint) (*domain.Conversation, error)
	CreateConversation(user1, user2 uint) (*domain.Conversation, error)
	SaveMessage(message *domain.Message) error
	FindAll(model interface{}, query string, args ...interface{}) error
	GetMessagesByConversation(convoID uint, limit, offset int) ([]domain.Message, error)
	FindConversationByID(id uint) (*domain.Conversation, error)
}
