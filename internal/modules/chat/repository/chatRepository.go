package repository

import (
	"gorm.io/gorm"
	domain "thinkdrop-backend/internal/Common"
	Domain "thinkdrop-backend/internal/modules/chat/domain"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) Domain.ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) FindConversation(user1, user2 uint) (*domain.Conversation, error) {
	var convo domain.Conversation

	err := r.db.Where(
		"(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)",
		user1, user2, user2, user1,
	).First(&convo).Error

	if err != nil {
		return nil, err
	}

	return &convo, nil
}

func (r *chatRepository) CreateConversation(user1, user2 uint) (*domain.Conversation, error) {
	convo := domain.Conversation{
		User1ID: user1,
		User2ID: user2,
	}

	if err := r.db.Create(&convo).Error; err != nil {
		return nil, err
	}

	return &convo, nil
}

func (r *chatRepository) SaveMessage(message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *chatRepository) FindAll(model interface{}, query string, args ...interface{}) error {
	return r.db.Where(query, args...).Find(model).Error
}
