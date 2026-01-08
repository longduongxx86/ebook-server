package handlers

import (
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/repositories"
	"main/internal/websocket"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationRepo *repositories.NotificationRepository
}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{
		notificationRepo: repositories.NewNotificationRepository(),
	}
}

// Helper to create and broadcast notification (từ original)
func createAndBroadcastNotification(title, message, typeStr string, referenceID uint) {
	notification := models.Notification{
		Title:       title,
		Message:     message,
		Type:        typeStr,
		ReferenceID: referenceID,
		IsRead:      false,
	}

	// Save to DB
	notificationRepo := repositories.NewNotificationRepository()
	if err := notificationRepo.Create(&notification); err != nil {
		// Log error but don't stop flow
		return
	}

	// Broadcast via WebSocket
	websocket.Manager.BroadcastNotification(title, message, typeStr, referenceID)
}

// GetNotifications returns a list of notifications (từ original)
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	// Only manager/admin
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access notifications"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// Get notifications
	notifications, total, err := h.notificationRepo.FindAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	// Count unread
	unreadCount, err := h.notificationRepo.GetUnreadCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get unread count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notifications": notifications,
		"unread_count":  unreadCount,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// MarkNotificationRead marks notifications as read (từ original)
func (h *NotificationHandler) MarkNotificationRead(c *gin.Context) {
	// Only manager/admin
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access notifications"})
		return
	}

	var req struct {
		IDs []uint `json:"ids"` // If empty, mark all as read
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mark as read
	if err := h.notificationRepo.MarkAsRead(req.IDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notifications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notifications marked as read"})
}