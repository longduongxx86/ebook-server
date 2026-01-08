package handlers

import (
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartRepo *repositories.CartRepository
	bookRepo *repositories.BookRepository
}

func NewCartHandler() *CartHandler {
	return &CartHandler{
		cartRepo: repositories.NewCartRepository(),
		bookRepo: repositories.NewBookRepository(),
	}
}

// AddToCart handles POST /cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var cartData struct {
		BookID   uint `json:"bookId" binding:"required"`
		Quantity int  `json:"quantity" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&cartData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check book stock
	hasStock, book, err := h.bookRepo.CheckStockAvailability(cartData.BookID, cartData.Quantity)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		}
		return
	}

	if !hasStock {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Get or create user's cart
	cart, err := h.cartRepo.GetOrCreateCart(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get or create cart"})
		return
	}

	// Check if book already exists in cart
	existingItem, err := h.cartRepo.GetCartItemByBook(cart.ID, cartData.BookID)
	if err == nil && existingItem != nil {
		// Update existing item quantity
		newQuantity := existingItem.Quantity + cartData.Quantity
		if newQuantity > book.Stock {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
			return
		}
		
		if err := h.cartRepo.UpdateCartItem(existingItem, newQuantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
			return
		}
	} else {
		// Create new cart item
		cartItem := &models.CartItem{
			CartID:   cart.ID,
			BookID:   cartData.BookID,
			Quantity: cartData.Quantity,
		}
		
		if err := h.cartRepo.CreateCartItem(cartItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
			return
		}
	}

	// Update cart timestamp
	if err := h.cartRepo.UpdateCartTimestamp(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

// GetCart handles GET /cart
func (h *CartHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	cart, err := h.cartRepo.GetCartByUserID(userID.(uint))
	if err != nil {
		if err.Error() == "record not found" {
			// Return empty cart
			c.JSON(http.StatusOK, gin.H{
				"cart": models.Cart{
					UserID: userID.(uint),
					Items:  []models.CartItem{},
				},
				"total_items": 0,
				"total_price": int64(0),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	// Calculate totals
	totalItems, totalPrice := h.cartRepo.CalculateCartTotals(cart)

	c.JSON(http.StatusOK, gin.H{
		"cart":        cart,
		"total_items": totalItems,
		"total_price": totalPrice,
	})
}

// UpdateCartItem handles PUT /cart/items/:item_id
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	itemID, err := strconv.ParseUint(c.Param("item_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updateData struct {
		Quantity int `json:"quantity" binding:"required,min=0"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get cart item
	cartItem, err := h.cartRepo.GetCartItemByID(uint(itemID), userID.(uint))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart item"})
		}
		return
	}

	// If quantity is 0, delete the item
	if updateData.Quantity == 0 {
		if err := h.cartRepo.DeleteCartItem(cartItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
		return
	}

	// Check stock availability
	hasStock, _, err := h.bookRepo.CheckStockAvailability(cartItem.BookID, updateData.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check stock"})
		return
	}
	
	if !hasStock {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Update quantity
	if err := h.cartRepo.UpdateCartItem(cartItem, updateData.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cart item updated successfully",
		"item":    cartItem,
	})
}

// RemoveFromCart handles DELETE /cart/items/:item_id
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	itemID, err := strconv.ParseUint(c.Param("item_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Get cart item
	cartItem, err := h.cartRepo.GetCartItemByID(uint(itemID), userID.(uint))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart item"})
		}
		return
	}

	// Delete cart item
	if err := h.cartRepo.DeleteCartItem(cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}

// ClearCart handles DELETE /cart
func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get user's cart
	cart, err := h.cartRepo.GetCartByUserID(userID.(uint))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		}
		return
	}

	// Delete all cart items
	if err := h.cartRepo.DeleteCartItemsByCartID(cart.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	// Update cart timestamp
	if err := h.cartRepo.UpdateCartTimestamp(cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
}

// GetCartSummary handles GET /cart/summary (lightweight)
func (h *CartHandler) GetCartSummary(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get cart items count
	count, err := h.cartRepo.GetCartItemsCount(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart summary"})
		return
	}

	// Get cart with minimal data for summary
	cart, err := h.cartRepo.GetCartByUserID(userID.(uint))
	var totalPrice int64 = 0
	
	if err == nil && cart != nil {
		totalItems, totalPriceCalc := h.cartRepo.CalculateCartTotals(cart)
		count = totalItems
		totalPrice = totalPriceCalc
	}

	c.JSON(http.StatusOK, gin.H{
		"item_count": count,
		"total_price": totalPrice,
	})
}