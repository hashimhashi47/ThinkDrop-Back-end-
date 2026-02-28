package repository

import (
	"log"
	domain "thinkdrop-backend/internal/Common"
	Domain "thinkdrop-backend/internal/modules/chat/domain"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) Domain.ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) GetAnonymousName(userID uint) (*domain.UserMini, error) {
	var user domain.UserMini

	err := r.db.
		Model(&domain.User{}).
		Select("anonymous_name,image_url").
		Where("id = ?", userID).
		Scan(&user).Error

	if err != nil {
		return nil, err
	}
	log.Println("data",user)

	return &user, nil
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

func (r *chatRepository) CreateConversation(user1, user2 uint, username1, user1image,
	username2, user2image string) (*domain.Conversation, error) {
	convo := domain.Conversation{
		User1ID:       user1,
		User2ID:       user2,
		User1NAME:     username1,
		User2NAME:     username2,
		User1ImageUrl: user1image,
		User2ImageUrl: user2image,
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

func (r *chatRepository) FindConversationByID(id uint) (*domain.Conversation, error) {
	var convo domain.Conversation
	if err := r.db.First(&convo, id).Error; err != nil {
		return nil, err
	}
	return &convo, nil
}

func (r *chatRepository) GetMessagesByConversation(convoID uint, limit, offset int) ([]domain.Message, error) {
	var messages []domain.Message

	err := r.db.
		Where("conversation_id = ?", convoID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error

	return messages, err
}
