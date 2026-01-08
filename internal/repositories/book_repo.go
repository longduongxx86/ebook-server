package repositories

import (
	"fmt"
	"main/internal/config"
	"main/internal/models"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		db: config.GetDB(),
	}
}

// Helper functions
func (r *BookRepository) slugify(s string) string {
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

// parseInt64FromQuery tries to parse an int64 from query param
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

// FindByID finds book by ID with category preloaded
func (r *BookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Category").First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// FindAllWithFilters finds books with pagination and filters
func (r *BookRepository) FindAllWithFilters(params BookQueryParams) (*BookQueryResult, error) {
	query := r.db.Model(&models.Book{})

	// Search filter
	if params.Search != "" {
		like := "%" + params.Search + "%"
		query = query.Where("title LIKE ? OR author LIKE ?", like, like)
	}

	// Price filters
	if params.MinPrice > 0 {
		query = query.Where("price >= ?", params.MinPrice)
	}
	if params.MaxPrice > 0 {
		query = query.Where("price <= ?", params.MaxPrice)
	}

	// Rating filter
	if params.MinRating > 0 {
		query = query.Where("average_rating >= ?", params.MinRating)
	}

	// Category filter
	if len(params.CategoryIDs) > 0 {
		query = query.Where("category_id IN ?", params.CategoryIDs)
	}

	// Stock filter (default: show only in stock)
	if params.ShowOutOfStock == false {
		query = query.Where("stock > 0")
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply sorting
	validSort := map[string]string{
		"created_at": "created_at",
		"price":      "price",
		"title":      "title",
		"rating":     "average_rating",
	}

	sortCol, ok := validSort[params.SortBy]
	if !ok {
		sortCol = "created_at"
	}

	sortOrder := "DESC"
	if params.SortOrder == "asc" {
		sortOrder = "ASC"
	}

	// Apply pagination
	offset := (params.Page - 1) * params.Limit
	query = query.Order(sortCol + " " + sortOrder).Offset(offset).Limit(params.Limit)

	// Execute query
	var books []models.Book
	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}

	// Load categories for mapping
	var categories []models.Category
	r.db.Find(&categories)
	catMap := make(map[uint]models.Category)
	for _, cat := range categories {
		catMap[cat.ID] = cat
	}

	// Calculate ratings and attach categories
	for i := range books {
		// Calculate rating
		var reviews []models.Review
		if err := r.db.Where("book_id = ?", books[i].ID).Find(&reviews).Error; err == nil {
			if len(reviews) > 0 {
				var sum int64
				for _, r := range reviews {
					sum += int64(r.Rating)
				}
				books[i].AverageRating = float64(sum) / float64(len(reviews))
				books[i].ReviewCount = len(reviews)
			} else {
				books[i].AverageRating = 0
				books[i].ReviewCount = 0
			}
		}

		// Attach category
		if books[i].CategoryID != nil {
			if cat, ok := catMap[*books[i].CategoryID]; ok {
				books[i].Category = &cat
			}
		}
	}

	return &BookQueryResult{
		Books: books,
		Total: total,
		Page:  params.Page,
		Limit: params.Limit,
	}, nil
}

// GetBookWithReviews returns a book with its reviews
func (r *BookRepository) GetBookWithReviews(id uint) (*models.Book, []models.Review, error) {
	var book models.Book
	err := r.db.Preload("Category").First(&book, id).Error
	if err != nil {
		return nil, nil, err
	}

	// Load reviews with user info
	var reviews []models.Review
	err = r.db.Where("book_id = ?", book.ID).
		Order("created_at DESC").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "full_name", "avatar_url")
		}).
		Find(&reviews).Error
	if err != nil {
		reviews = []models.Review{}
	}

	// Calculate average rating
	if len(reviews) > 0 {
		var sum int64
		for _, r := range reviews {
			sum += int64(r.Rating)
		}
		book.AverageRating = float64(sum) / float64(len(reviews))
		book.ReviewCount = len(reviews)
	}

	return &book, reviews, nil
}

// Create creates a new book
func (r *BookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

// Update updates a book
func (r *BookRepository) Update(id uint, patch map[string]interface{}) (*models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&book).Updates(patch).Error; err != nil {
		return nil, err
	}

	// Return updated book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

// Delete deletes a book and its related cart items
func (r *BookRepository) Delete(id uint) error {
	// Check if book has orders
	var orderCount int64
	if err := r.db.Model(&models.OrderItem{}).Where("book_id = ?", id).Count(&orderCount).Error; err != nil {
		return err
	}
	if orderCount > 0 {
		return fmt.Errorf("cannot delete book because it belongs to one or more orders")
	}

	// Start transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete related cart items
	if err := tx.Where("book_id = ?", id).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete book
	if err := tx.Delete(&models.Book{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// CountOrderItems counts order items for a book
func (r *BookRepository) CountOrderItems(bookID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.OrderItem{}).Where("book_id = ?", bookID).Count(&count).Error
	return count, err
}


// FindByIDWithPreload finds book by ID with category preloaded
func (r *BookRepository) FindByIDWithPreload(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Category").First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) CheckStockAvailability(bookID uint, quantity int) (bool, *models.Book, error) {
	var book models.Book
	err := r.db.First(&book, bookID).Error
	if err != nil {
		return false, nil, err
	}
	return book.Stock >= quantity, &book, nil
}

// UpdateStock updates book stock (for order processing)
func (r *BookRepository) UpdateStock(bookID uint, quantity int, isDecrement bool) error {
	var book models.Book
	if err := r.db.First(&book, bookID).Error; err != nil {
		return err
	}
	
	if isDecrement {
		book.Stock -= quantity
	} else {
		book.Stock += quantity
	}
	
	return r.db.Save(&book).Error
}

// DTOs for repository
type BookQueryParams struct {
	Page           int
	Limit          int
	Search         string
	MinPrice       int64
	MaxPrice       int64
	MinRating      float64
	CategoryIDs    []uint
	SortBy         string
	SortOrder      string
	ShowOutOfStock bool
}

type BookQueryResult struct {
	Books []models.Book
	Total int64
	Page  int
	Limit int
}
