package repositories

import (
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{
		db: config.GetDB(),
	}
}

// Create creates a new notification (chỉ hàm này)
func (r *NotificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

// FindAll gets notifications with pagination (chỉ hàm này)
func (r *NotificationRepository) FindAll(page, limit int) ([]models.Notification, int64, error) {
	offset := (page - 1) * limit

	var notifications []models.Notification
	var total int64

	// Get total count
	if err := r.db.Model(&models.Notification{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get notifications
	err := r.db.Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&notifications).Error

	return notifications, total, err
}

// MarkAsRead marks notifications as read (chỉ hàm này)
func (r *NotificationRepository) MarkAsRead(ids []uint) error {
	query := r.db.Model(&models.Notification{}).Where("is_read = ?", false)

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}

	return query.Updates(map[string]interface{}{
		"is_read":    true,
		"updated_at": time.Now().UnixMilli(),
	}).Error
}

// GetUnreadCount gets count of unread notifications (chỉ hàm này)
func (r *NotificationRepository) GetUnreadCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Notification{}).
		Where("is_read = ?", false).
		Count(&count).Error
	return count, err
}