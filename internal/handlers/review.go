package handlers

import (
	"net/http"
	"strconv"
	"time"

	"main/internal/models"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewRepo *repositories.ReviewRepository
	bookRepo   *repositories.BookRepository
}

func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{
		reviewRepo: repositories.NewReviewRepository(),
		bookRepo:   repositories.NewBookRepository(),
	}
}

// CreateReview handles POST /reviews
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var reviewData struct {
		Rating  int    `json:"rating" binding:"required,min=1,max=5"`
		Comment string `json:"comment"`
		BookID  int64  `json:"bookId"`
	}

	if err := c.ShouldBindJSON(&reviewData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if book exists
	_, err := h.bookRepo.FindByID(uint(reviewData.BookID))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		}
		return
	}

	// Check if user already reviewed this book
	existingReview, err := h.reviewRepo.FindByUserAndBook(
		userID.(uint),
		uint(reviewData.BookID),
	)
	if err == nil && existingReview != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Bạn đã đánh giá cuốn sách này rồi!"})
		return
	}

	// Create review
	review := &models.Review{
		BookID:  uint(reviewData.BookID),
		UserID:  userID.(uint),
		Rating:  reviewData.Rating,
		Comment: reviewData.Comment,
	}

	if err := h.reviewRepo.Create(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	// Update book's average rating
	if err := h.reviewRepo.UpdateBookRating(uint(reviewData.BookID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book rating"})
		return
	}

	// Load review with user info
	if err := h.reviewRepo.GetReviewsWithUser(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review created successfully",
		"review":  review,
	})
}

// UpdateReview handles PUT /reviews/:review_id
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	reviewID, err := strconv.ParseUint(c.Param("review_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	var updateData struct {
		Rating  int    `json:"rating" binding:"required,min=1,max=5"`
		Comment string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find review
	review, err := h.reviewRepo.FindByUserAndReviewID(
		userID.(uint),
		uint(reviewID),
	)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		}
		return
	}

	// Prepare updates
	updates := map[string]interface{}{
		"rating":     updateData.Rating,
		"comment":    updateData.Comment,
		"updated_at": time.Now().UnixMilli(),
	}

	// Update review
	if err := h.reviewRepo.Update(review, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		return
	}

	// Update book's average rating
	if err := h.reviewRepo.UpdateBookRating(review.BookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book rating"})
		return
	}

	// Load review with user info
	if err := h.reviewRepo.GetReviewsWithUser(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Review updated successfully",
		"review":  review,
	})
}

// DeleteReview handles DELETE /reviews/:review_id
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	reviewID, err := strconv.ParseUint(c.Param("review_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	// Find review
	review, err := h.reviewRepo.FindByUserAndReviewID(
		userID.(uint),
		uint(reviewID),
	)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		}
		return
	}

	bookID := review.BookID

	// Delete review
	if err := h.reviewRepo.Delete(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	// Update book's average rating
	if err := h.reviewRepo.UpdateBookRating(bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book rating"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

// GetBookReviews handles GET /books/:id/reviews
func (h *ReviewHandler) GetBookReviews(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Parse pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Check if user is admin to see unapproved reviews
	approvedOnly := true
	userRole, exists := c.Get("user_role")
	if exists && (userRole == "admin" || userRole == "manager") {
		// Admin/manager can see all reviews
		approvedOnly = false
	}

	// Get reviews
	reviews, total, err := h.reviewRepo.FindByBookID(
		uint(bookID),
		page,
		limit,
		approvedOnly,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reviews": reviews,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// AdminApproveReview (thêm nếu cần)
func (h *ReviewHandler) AdminApproveReview(c *gin.Context) {
	// Check admin role
	userRole, exists := c.Get("user_role")
	if !exists || (userRole != "admin" && userRole != "manager") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can approve reviews"})
		return
	}

	reviewID, err := strconv.ParseUint(c.Param("review_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	review, err := h.reviewRepo.FindByID(uint(reviewID))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		}
		return
	}

	// Approve review
	updates := map[string]interface{}{
		"approved":   true,
		"updated_at": time.Now().UnixMilli(),
	}

	if err := h.reviewRepo.Update(review, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review approved successfully"})
}