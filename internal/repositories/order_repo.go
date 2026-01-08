package repositories

import (
	"fmt"
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: config.GetDB(),
	}
}

// CreateOrder creates a new order with transaction
func (r *OrderRepository) CreateOrder(order *models.Order, items []struct {
	BookID   uint
	Quantity int
}) (*models.Order, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Generate order number
	order.OrderNumber = r.generateOrderNumber()
	order.Status = "pending"
	order.CreatedAt = time.Now().UnixMilli()
	order.UpdatedAt = time.Now().UnixMilli()

	// Create order
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Process items
	var totalAmount int64
	for _, item := range items {
		// Check book stock
		var book models.Book
		if err := tx.First(&book, item.BookID).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("book %d not found: %w", item.BookID, err)
		}

		if book.Stock < item.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("insufficient stock for book %s", book.Title)
		}

		// Create order item
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			BookID:    item.BookID,
			Quantity:  item.Quantity,
			Price:     book.Price,
			Cost:      book.Cost,
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}

		// Update book stock
		book.Stock -= item.Quantity
		if err := tx.Save(&book).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update book stock: %w", err)
		}

		totalAmount += book.Price * int64(item.Quantity)
	}

	// Update order total
	order.TotalAmount = totalAmount
	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update order total: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return order, nil
}

// CreateOrderFromCart creates order from cart
func (r *OrderRepository) CreateOrderFromCart(userID uint, shippingAddress, notes string) (*models.Order, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get cart with items
	var cart models.Cart
	if err := tx.Preload("Items.Book").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("cart not found: %w", err)
	}

	if len(cart.Items) == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("cart is empty")
	}

	// Create order
	order := models.Order{
		OrderNumber:     r.generateOrderNumber(),
		BuyerID:         userID,
		TotalAmount:     0,
		Status:          "pending",
		ShippingAddress: shippingAddress,
		Notes:           notes,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Process cart items
	var totalAmount int64
	for _, ci := range cart.Items {
		// Re-fetch book for latest stock
		var book models.Book
		if err := tx.First(&book, ci.BookID).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("book %d not found: %w", ci.BookID, err)
		}

		if book.Stock < ci.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("insufficient stock for book %s", book.Title)
		}

		// Create order item
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			BookID:    ci.BookID,
			Quantity:  ci.Quantity,
			Price:     book.Price,
			Cost:      book.Cost,
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}

		// Update book stock
		book.Stock -= ci.Quantity
		if err := tx.Save(&book).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update book stock: %w", err)
		}

		totalAmount += book.Price * int64(ci.Quantity)
	}

	// Update order total
	order.TotalAmount = totalAmount
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update order total: %w", err)
	}

	// Clear cart
	if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to clear cart: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &order, nil
}

// FindByUserID finds orders by user with pagination and filters
func (r *OrderRepository) FindByUserID(userID uint, page, limit int, status string) ([]models.Order, int64, error) {
	offset := (page - 1) * limit

	query := r.db.Where("buyer_id = ?", userID)

	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	var total int64
	if err := query.Model(&models.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get orders
	var orders []models.Order
	err := query.Preload("Items.Book").Preload("Buyer").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, total, err
}

// FindAll finds all orders (for managers/admin)
func (r *OrderRepository) FindAll(page, limit int, status string) ([]models.Order, int64, error) {
	offset := (page - 1) * limit

	query := r.db.Model(&models.Order{})

	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get orders
	var orders []models.Order
	err := query.Preload("Items.Book").Preload("Buyer").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error

	return orders, total, err
}

// FindByID finds order by ID
func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items.Book").Preload("Buyer").
		First(&order, id).Error
	return &order, err
}

// FindByIDAndUser finds order by ID and user ID
func (r *OrderRepository) FindByIDAndUser(id, userID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items.Book").Preload("Buyer").
		Where("id = ? AND buyer_id = ?", id, userID).
		First(&order).Error
	return &order, err
}

// UpdateStatus updates order status
func (r *OrderRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now().UnixMilli(),
		}).Error
}

// CancelOrder cancels order and restores stock
func (r *OrderRepository) CancelOrder(orderID, userID uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get order with items
	var order models.Order
	if err := tx.Preload("Items").Where("id = ? AND buyer_id = ?", orderID, userID).First(&order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("order not found: %w", err)
	}

	// Check if can be cancelled
	if order.Status != "pending" {
		tx.Rollback()
		return fmt.Errorf("cannot cancel order with status: %s", order.Status)
	}

	// Restore book stock
	for _, item := range order.Items {
		if err := tx.Model(&models.Book{}).
			Where("id = ?", item.BookID).
			Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to restore stock: %w", err)
		}
	}

	// Update order status
	if err := tx.Model(&order).Updates(map[string]interface{}{
		"status":     "cancelled",
		"updated_at": time.Now().UnixMilli(),
	}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to cancel order: %w", err)
	}

	return tx.Commit().Error
}

// GetUserOrdersCount gets count of orders for user
func (r *OrderRepository) GetUserOrdersCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Order{}).
		Where("buyer_id = ?", userID).
		Count(&count).Error
	return count, err
}

// Helper function
func (r *OrderRepository) generateOrderNumber() string {
	return fmt.Sprintf("ORD-%d", time.Now().Unix())
}
func (r *OrderRepository) GetOrderWithBuyer(orderID, userID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Where("id = ? AND buyer_id = ?", orderID, userID).First(&order).Error
	return &order, err
}