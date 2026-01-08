package repositories

import (
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		db: config.GetDB(),
	}
}

// CreateCart creates a new cart for user
func (r *CartRepository) CreateCart(userID uint) (*models.Cart, error) {
	cart := &models.Cart{
		UserID: userID,
	}
	err := r.db.Create(cart).Error
	return cart, err
}

// GetCartByUserID gets cart by user ID with items and books preloaded
func (r *CartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items.Book").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// GetOrCreateCart gets or creates cart for user
func (r *CartRepository) GetOrCreateCart(userID uint) (*models.Cart, error) {
	cart, err := r.GetCartByUserID(userID)
	if err == nil {
		return cart, nil
	}
	
	// If cart not found, create new one
	if err.Error() == "record not found" {
		return r.CreateCart(userID)
	}
	
	return nil, err
}

// GetCartItemByID gets cart item by ID with book info
func (r *CartRepository) GetCartItemByID(itemID uint, userID uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Preload("Book").Preload("Cart").
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("cart_items.id = ? AND carts.user_id = ?", itemID, userID).
		First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

// GetCartItemByBook gets cart item by book ID in user's cart
func (r *CartRepository) GetCartItemByBook(cartID uint, bookID uint) (*models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.Where("cart_id = ? AND book_id = ?", cartID, bookID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

// CreateCartItem creates a new cart item
func (r *CartRepository) CreateCartItem(cartItem *models.CartItem) error {
	return r.db.Create(cartItem).Error
}

// UpdateCartItem updates cart item quantity
func (r *CartRepository) UpdateCartItem(cartItem *models.CartItem, quantity int) error {
	cartItem.Quantity = quantity
	cartItem.UpdatedAt = time.Now().UnixMilli()
	return r.db.Save(cartItem).Error
}

// DeleteCartItem deletes a cart item
func (r *CartRepository) DeleteCartItem(cartItem *models.CartItem) error {
	return r.db.Delete(cartItem).Error
}

// DeleteCartItemsByCartID deletes all items from cart
func (r *CartRepository) DeleteCartItemsByCartID(cartID uint) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}

// UpdateCartTimestamp updates cart's updated_at timestamp
func (r *CartRepository) UpdateCartTimestamp(cart *models.Cart) error {
	cart.UpdatedAt = time.Now().UnixMilli()
	return r.db.Save(cart).Error
}

// CalculateCartTotals calculates total items and price in cart
func (r *CartRepository) CalculateCartTotals(cart *models.Cart) (int, int64) {
	var totalItems int
	var totalPrice int64
	
	for _, item := range cart.Items {
		totalItems += item.Quantity
		if item.Book.ID != 0 {
			totalPrice += item.Book.Price * int64(item.Quantity)
		}
	}
	
	return totalItems, totalPrice
}

// GetCartItemsCount gets count of items in user's cart
func (r *CartRepository) GetCartItemsCount(userID uint) (int, error) {
	var count int64
	err := r.db.Model(&models.CartItem{}).
		Joins("JOIN carts ON carts.id = cart_items.cart_id").
		Where("carts.user_id = ?", userID).
		Count(&count).Error
	return int(count), err
}