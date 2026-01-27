package handlers

import (
	"main/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatRepo *repositories.ChatRepository
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		chatRepo: repositories.NewChatRepository(),
	}
}

// GetConversations returns a list of conversations (Admin only)
func (h *ChatHandler) GetConversations(c *gin.Context) {
	conversations, err := h.chatRepo.GetConversations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, conversations)
}

// GetChatHistory returns messages for a conversation
func (h *ChatHandler) GetChatHistory(c *gin.Context) {
	userIDVal, _ := c.Get("user_id")
	roleVal, _ := c.Get("user_role")

	userID := userIDVal.(uint)
	role := roleVal.(string)

	var targetUserID uint

	if role == "admin" || role == "manager" {
		targetUserIDStr := c.Query("user_id")
		if targetUserIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required for admin"})
			return
		}
		id, err := strconv.Atoi(targetUserIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
			return
		}
		targetUserID = uint(id)
	} else {
		targetUserID = userID // User can only see their own chat
	}

	conversation, err := h.chatRepo.GetConversationByUserID(targetUserID)
	if err != nil {
		// Return empty list if no conversation yet
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	messages, err := h.chatRepo.GetMessagesByConversationID(conversation.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
