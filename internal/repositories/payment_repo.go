package repositories

import (
	"fmt"
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		db: config.GetDB(),
	}
}

// Create creates a new payment
func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

// FindByOrderID finds payment by order ID
func (r *PaymentRepository) FindByOrderID(orderID uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

// FindByID finds payment by ID
func (r *PaymentRepository) FindByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order").First(&payment, id).Error
	return &payment, err
}

// FindAll gets all payments (for admin)
func (r *PaymentRepository) FindAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Preload("Order").Preload("Order.Buyer").Order("created_at DESC").Find(&payments).Error
	return payments, err
}

// FindByUserID gets payments for a specific user
func (r *PaymentRepository) FindByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Table("payments").
		Joins("JOIN orders ON orders.id = payments.order_id").
		Where("orders.buyer_id = ?", userID).
		Preload("Order").
		Order("payments.created_at DESC").
		Find(&payments).Error
	return payments, err
}

// UpdateStatus updates payment status
func (r *PaymentRepository) UpdateStatus(id uint, status, transactionID string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.Preload("Order").First(&payment, id).Error; err != nil {
		return nil, err
	}

	payment.Status = status
	if transactionID != "" {
		payment.TransactionID = transactionID
	}
	payment.UpdatedAt = time.Now().UnixMilli()

	if err := r.db.Save(&payment).Error; err != nil {
		return nil, err
	}

	// If payment completed, update order status
	if status == "completed" {
		payment.Order.Status = "confirmed"
		payment.Order.UpdatedAt = time.Now().UnixMilli()
		if err := r.db.Save(&payment.Order).Error; err != nil {
			return nil, err
		}
	}

	return &payment, nil
}

// CheckOrderExists checks if order exists and belongs to user
func (r *PaymentRepository) CheckOrderExists(orderID, userID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Where("id = ? AND buyer_id = ?", orderID, userID).First(&order).Error
	return &order, err
}

// CheckPaymentExists checks if payment already exists for order
func (r *PaymentRepository) CheckPaymentExists(orderID uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

// Generate mock QR code data
func (r *PaymentRepository) GenerateQRCode(orderNumber string, amount int64) string {
	return fmt.Sprintf("QR:%s:%d:%d", orderNumber, amount, time.Now().UnixMilli())
}

// Generate mock bank info
func (r *PaymentRepository) GenerateBankInfo() string {
	return `{
		"bank_name": "Vietcombank",
		"account_number": "1234567890",
		"account_name": "EBook Store",
		"branch": "Ho Chi Minh City"
	}`
}