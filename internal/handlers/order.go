package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orderRepo: repositories.NewOrderRepository(),
	}
}

// CreateOrder handles POST /orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var orderData struct {
		Items []struct {
			BookID   uint `json:"book_id" binding:"required"`
			Quantity int  `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
		ShippingAddress string `json:"shipping_address" binding:"required"`
		Notes           string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tạo order object
	order := &models.Order{
		BuyerID:         userID.(uint),
		ShippingAddress: orderData.ShippingAddress,
		Notes:           orderData.Notes,
	}

	// Gọi repository (trả về *models.Order)
	order, err := h.orderRepo.CreateOrder(order, []struct {
		BookID   uint
		Quantity int
	}(orderData.Items))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Load full order details
	order, err = h.orderRepo.FindByID(order.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
		return
	}

	// Send notification
	go func() {
		createAndBroadcastNotification(
			"Đơn hàng mới",
			fmt.Sprintf("Đơn hàng mới %s đã được tạo", order.OrderNumber),
			"order",
			order.ID,
		)
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// GetOrders handles GET /orders
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	role, _ := c.Get("user_role")

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")

	var orders []models.Order
	var total int64
	var err error

	if role == "manager" || role == "admin" {
		// Admin can see all orders
		orders, total, err = h.orderRepo.FindAll(page, limit, status)
	} else {
		// Regular user sees only their orders
		orders, total, err = h.orderRepo.FindByUserID(userID.(uint), page, limit, status)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
		"filter": gin.H{
			"status": status,
		},
	})
}

// GetOrder handles GET /orders/:id
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	role, _ := c.Get("user_role")

	var order *models.Order

	if role == "manager" || role == "admin" {
		// Admin can see any order
		order, err = h.orderRepo.FindByID(uint(id))
	} else {
		// Regular user can only see their own orders
		order, err = h.orderRepo.FindByIDAndUser(uint(id), userID.(uint))
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// CreateOrderFromCart handles POST /orders/from-cart
func (h *OrderHandler) CreateOrderFromCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var body struct {
		ShippingAddress string `json:"shipping_address" binding:"required"`
		Notes           string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create order from cart
	order, err := h.orderRepo.CreateOrderFromCart(userID.(uint), body.ShippingAddress, body.Notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Load full order details
	order, err = h.orderRepo.FindByID(order.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
		return
	}

	// Send notification
	go func() {
		createAndBroadcastNotification(
			"Đơn hàng mới",
			fmt.Sprintf("Đơn hàng mới %s đã được tạo", order.OrderNumber),
			"order",
			order.ID,
		)
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// CancelOrder handles PUT /orders/:id/cancel
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.orderRepo.CancelOrder(uint(orderID), userID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get updated order
	order, err := h.orderRepo.FindByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order cancelled successfully",
		"order":   order,
	})
}

// UpdateOrderStatus handles PUT /orders/:id/status (admin only)
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	// Check admin role
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can update order status"})
		return
	}

	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var statusData struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate status
	validStatuses := []string{"pending", "confirmed", "shipped", "delivered", "cancelled"}
	if !contains(validStatuses, statusData.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	// Update status
	err = h.orderRepo.UpdateStatus(uint(orderID), statusData.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	// Get updated order
	order, err := h.orderRepo.FindByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
		"order":   order,
	})
}

// GetOrderStats handles GET /orders/stats (admin only)
func (h *OrderHandler) GetOrderStats(c *gin.Context) {
	// Check admin role
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can view order stats"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get user's order count
	count, err := h.orderRepo.GetUserOrdersCount(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order stats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_count": count,
	})
}
