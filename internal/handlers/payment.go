package handlers

import (
	"net/http"
	"strconv"

	"main/internal/models"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentRepo *repositories.PaymentRepository
	orderRepo   *repositories.OrderRepository
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		paymentRepo: repositories.NewPaymentRepository(),
		orderRepo:   repositories.NewOrderRepository(),
	}
}

// CreatePayment handles POST /orders/:id/payments
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
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

	var paymentData struct {
		Method string `json:"method" binding:"required"` // qr, cash, bank_transfer
	}

	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate payment method
	validMethods := []string{"qr", "cash", "bank_transfer"}
	if !contains(validMethods, paymentData.Method) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment method"})
		return
	}

	// Check if order exists and belongs to user
	order, err := h.paymentRepo.CheckOrderExists(uint(orderID), userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Check if payment already exists
	if _, err := h.paymentRepo.CheckPaymentExists(uint(orderID)); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Payment already exists for this order"})
		return
	}

	// Create payment
	payment := &models.Payment{
		OrderID:   uint(orderID),
		Amount:    int64(order.TotalAmount),
		Method:    paymentData.Method,
		Status:    "pending",
	}

	// Generate QR code and bank info for QR payment
	if paymentData.Method == "qr" {
		payment.QRCode = h.paymentRepo.GenerateQRCode(order.OrderNumber, int64(order.TotalAmount))
		payment.BankInfo = h.paymentRepo.GenerateBankInfo()
	}

	if err := h.paymentRepo.Create(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment created successfully",
		"payment": payment,
	})
}

// GetPayment handles GET /orders/:id/payment
func (h *PaymentHandler) GetPayment(c *gin.Context) {
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

	// Check if order belongs to user
	if _, err := h.paymentRepo.CheckOrderExists(uint(orderID), userID.(uint)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Get payment
	payment, err := h.paymentRepo.FindByOrderID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

// GetAllPayments handles GET /payments (Manager only)
func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access payments"})
		return
	}

	payments, err := h.paymentRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

// GetUserPayments handles GET /payments/my-payments
func (h *PaymentHandler) GetUserPayments(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	payments, err := h.paymentRepo.FindByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

// UpdatePaymentStatus handles PUT /payments/:payment_id/status
func (h *PaymentHandler) UpdatePaymentStatus(c *gin.Context) {
	paymentID, err := strconv.ParseUint(c.Param("payment_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	var statusData struct {
		Status        string `json:"status" binding:"required"`
		TransactionID string `json:"transaction_id"`
	}

	if err := c.ShouldBindJSON(&statusData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate status
	validStatuses := []string{"pending", "completed", "failed"}
	if !contains(validStatuses, statusData.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	// Update payment status
	payment, err := h.paymentRepo.UpdateStatus(uint(paymentID), statusData.Status, statusData.TransactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment status updated successfully",
		"payment": payment,
	})
}

// Helper function (từ original)
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}