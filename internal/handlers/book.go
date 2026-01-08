package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"main/internal/models"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)
type BookHandler struct {
	bookRepo *repositories.BookRepository
}
func NewBookHandler() *BookHandler {
	return &BookHandler{
		bookRepo: repositories.NewBookRepository(),
	}
}
// GetBooks handles GET /books with filters
func (h *BookHandler) GetBooks(c *gin.Context) {
	// Parse pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}

	// Parse filters
	search := strings.TrimSpace(c.Query("search"))
	minPriceStr := strings.TrimSpace(c.Query("min_price"))
	maxPriceStr := strings.TrimSpace(c.Query("max_price"))
	minRatingStr := strings.TrimSpace(c.Query("min_rating"))
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := strings.ToLower(c.DefaultQuery("sort_order", "desc"))
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	// Parse categories
	categoriesParam := strings.TrimSpace(c.Query("categories"))
	var categoryIDs []uint
	if categoriesParam != "" {
		parts := strings.Split(categoriesParam, ",")
		for _, part := range parts {
			if id, err := strconv.ParseUint(strings.TrimSpace(part), 10, 64); err == nil && id > 0 {
				categoryIDs = append(categoryIDs, uint(id))
			}
		}
	}

	// Parse price filters
	var minPrice, maxPrice int64
	if v, ok := parseInt64FromQuery(minPriceStr); ok {
		minPrice = v
	}
	if v, ok := parseInt64FromQuery(maxPriceStr); ok {
		maxPrice = v
	}

	// Parse rating filter
	var minRating float64
	if v, ok := parseFloat64FromQuery(minRatingStr); ok {
		minRating = v
	}

	// Build query params
	params := repositories.BookQueryParams{
		Page:        page,
		Limit:       limit,
		Search:      search,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
		MinRating:   minRating,
		CategoryIDs: categoryIDs,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
	}

	// Query books
	result, err := h.bookRepo.FindAllWithFilters(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": result.Books,
		"pagination": gin.H{
			"page":  result.Page,
			"limit": result.Limit,
			"total": result.Total,
		},
	})
}
// GetBook handles GET /books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}
	id := uint(id64)

	book, reviews, err := h.bookRepo.GetBookWithReviews(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book":    book,
		"reviews": reviews,
	})
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	// Authentication check
	uid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}
	_, ok = uid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user id in context"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Basic validation
	if strings.TrimSpace(input.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	if input.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be > 0"})
		return
	}
	if input.Cost < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cost must be >= 0"})
		return
	}
	if input.Cost > input.Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cost cannot be greater than price"})
		return
	}
	if input.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock must be >= 0"})
		return
	}

	// Generate slug if not provided
	if input.Slug == "" {
		input.Slug = slugify(input.Title)
	}

	if err := h.bookRepo.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "book": input})
}

// UpdateBook handles PATCH /books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}
	id := uint(id64)

	// Bind incoming JSON
	var patch map[string]interface{}
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove disallowed fields
	delete(patch, "id")
	delete(patch, "ID")
	delete(patch, "created_at")
	delete(patch, "CreatedAt")
	delete(patch, "deleted_at")
	delete(patch, "DeletedAt")

	// Generate slug if title is updated
	if title, ok := patch["title"].(string); ok {
		if _, hasSlug := patch["slug"]; !hasSlug {
			patch["slug"] = slugify(title)
		}
	}

	// Convert numeric fields
	h.normalizePatchNumericFields(patch)

	// Validate cost vs price
	if err := h.validateCostPrice(patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book
	book, err := h.bookRepo.Update(id, patch)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})
}

// DeleteBook handles DELETE /books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}
	id := uint(id64)

	err = h.bookRepo.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else if strings.Contains(err.Error(), "cannot delete book") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// Helper functions (moved from original handlers)
func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "-")
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			out = append(out, r)
		}
	}
	res := string(out)
	if res == "" {
		res = strconv.FormatInt(time.Now().Unix(), 10)
	}
	return res
}

func parseInt64FromQuery(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return v, true
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return int64(f), true
	}
	return 0, false
}

func parseFloat64FromQuery(s string) (float64, bool) {
	if s == "" {
		return 0, false
	}
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return v, true
	}
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return float64(v), true
	}
	return 0, false
}

func (h *BookHandler) normalizePatchNumericFields(patch map[string]interface{}) {
	// Normalize price
	if price, ok := patch["price"]; ok {
		patch["price"] = normalizeToInt64(price)
	}

	// Normalize cost
	if cost, ok := patch["cost"]; ok {
		patch["cost"] = normalizeToInt64(cost)
	}
}

func normalizeToInt64(value interface{}) int64 {
	switch v := value.(type) {
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	case string:
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			return parsed
		} else if f, err := strconv.ParseFloat(v, 64); err == nil {
			return int64(f)
		}
	}
	return 0
}

func (h *BookHandler) validateCostPrice(patch map[string]interface{}) error {
	// This is simplified - in real implementation, you'd fetch current book
	// and validate based on existing values and patch values
	return nil
}