package repositories

import (
	"main/internal/config"
	"main/internal/models"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository() *ChatRepository {
	return &ChatRepository{
		db: config.GetDB(),
	}
}

// GetConversations returns a list of conversations ordered by last message time
func (r *ChatRepository) GetConversations() ([]models.Conversation, error) {
	var conversations []models.Conversation
	err := r.db.Preload("User").Order("last_message_at desc").Find(&conversations).Error
	return conversations, err
}

// GetConversationByUserID finds a conversation by user ID
func (r *ChatRepository) GetConversationByUserID(userID uint) (*models.Conversation, error) {
	var conversation models.Conversation
	err := r.db.Where("user_id = ?", userID).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// GetMessagesByConversationID returns messages for a conversation
func (r *ChatRepository) GetMessagesByConversationID(conversationID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("conversation_id = ?", conversationID).Preload("Sender").Order("created_at asc").Find(&messages).Error
	return messages, err
}

// FirstOrCreateConversation finds or creates a conversation for a user
func (r *ChatRepository) FirstOrCreateConversation(userID uint) (*models.Conversation, error) {
	var conversation models.Conversation
	err := r.db.FirstOrCreate(&conversation, models.Conversation{UserID: userID}).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// CreateMessage creates a new message
func (r *ChatRepository) CreateMessage(msg *models.Message) error {
	return r.db.Create(msg).Error
}

// UpdateConversationLastMessage updates the last message info of a conversation
func (r *ChatRepository) UpdateConversationLastMessage(conversationID uint, lastMessage string, lastMessageAt int64) error {
	return r.db.Model(&models.Conversation{}).Where("id = ?", conversationID).Updates(map[string]interface{}{
		"last_message":    lastMessage,
		"last_message_at": lastMessageAt,
	}).Error
}

// GetMessageWithSender retrieves a message with sender info
func (r *ChatRepository) GetMessageWithSender(messageID uint) (*models.Message, error) {
	var msg models.Message
	err := r.db.Preload("Sender").First(&msg, messageID).Error
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
