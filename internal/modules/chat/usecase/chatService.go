package usecase

import (
	"errors"
	domain "thinkdrop-backend/internal/Common"
	Domain "thinkdrop-backend/internal/modules/chat/domain"
	"time"
)

type ChatService struct {
	repo Domain.ChatRepository
}

func NewChatService(repo Domain.ChatRepository) *ChatService {
	return &ChatService{repo}
}

func (s *ChatService) SendMessage(senderID, receiverID uint, content string) (*domain.Message, error) {

	// check conversation
	convo, err := s.repo.FindConversation(senderID, receiverID)

	if err != nil {
		// if not found, create new
		convo, err = s.repo.CreateConversation(senderID, receiverID)
		if err != nil {
			return nil, err
		}
	}

	message := &domain.Message{
		ConversationID: convo.ID,
		SenderID:       senderID,
		ReceiverID:     receiverID,
		Content:        content,
		CreatedAt:      time.Now(),
	}

	if err := s.repo.SaveMessage(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *ChatService) Getallchat(userID uint) ([]domain.Conversation, error) {
	var Conversation []domain.Conversation

	if err := s.repo.FindAll(&Conversation, "user1_id = ? OR user2_id = ?", userID, userID); err != nil {
		return nil, err
	}

	var Users []domain.User

	for _, v := range Conversation {
		if v.User1ID != userID {
			if err := s.repo.FindAll(&Users, "id = ?", v.User1ID); err != nil {
				return nil, errors.New("failed to find the users")
			}
		}
		if v.User2ID != userID {
			if err := s.repo.FindAll(&Users, "id = ?", v.User2ID); err != nil {
				return nil, errors.New("failed to find the users")
			}
		}
	}
	
	return Conversation, nil
}

func (s *ChatService) GetMessages(userID, convoID uint, limit, offset int) ([]domain.Message, error) {

	// optional security check
	convo, err := s.repo.FindConversationByID(convoID)
	if err != nil {
		return nil, err
	}

	if convo.User1ID != userID && convo.User2ID != userID {
		return nil, errors.New("unauthorized")
	}

	return s.repo.GetMessagesByConversation(convoID, limit, offset)
}

func (s *ChatService) StartConversation(user1, user2 uint) (*domain.Conversation, error) {

	convo, err := s.repo.FindConversation(user1, user2)
	if err == nil {
		return convo, nil
	}

	return s.repo.CreateConversation(user1, user2)
}
