package repositories

import (
	"main/internal/config"
	"main/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		db: config.GetDB(),
	}
}

// Create creates a new review
func (r *ReviewRepository) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

// FindByID finds review by ID
func (r *ReviewRepository) FindByID(id uint) (*models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// FindByUserAndBook finds review by user and book
func (r *ReviewRepository) FindByUserAndBook(userID, bookID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("book_id = ? AND user_id = ?", bookID, userID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// FindByUserAndReviewID finds review by user and review ID
func (r *ReviewRepository) FindByUserAndReviewID(userID, reviewID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.Where("id = ? AND user_id = ?", reviewID, userID).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// Update updates a review
func (r *ReviewRepository) Update(review *models.Review, updates map[string]interface{}) error {
	return r.db.Model(review).Updates(updates).Error
}

// Delete deletes a review
func (r *ReviewRepository) Delete(review *models.Review) error {
	return r.db.Delete(review).Error
}

// FindByBookID finds reviews for a book with pagination
func (r *ReviewRepository) FindByBookID(bookID uint, page, limit int, approvedOnly bool) ([]models.Review, int64, error) {
	offset := (page - 1) * limit

	// Build query
	query := r.db.Where("book_id = ?", bookID)
	if approvedOnly {
		query = query.Where("approved = ?", true)
	}

	// Get total count
	var total int64
	if err := query.Model(&models.Review{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get reviews with user preloaded
	var reviews []models.Review
	query = query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "full_name", "avatar_url")
	}).Order("created_at DESC").
		Offset(offset).
		Limit(limit)

	if err := query.Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

// UpdateBookRating updates book's average rating
func (r *ReviewRepository) UpdateBookRating(bookID uint) error {
	var reviews []models.Review
	if err := r.db.Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
		return err
	}

	var book models.Book
	if err := r.db.First(&book, bookID).Error; err != nil {
		return err
	}

	if len(reviews) > 0 {
		var totalRating float64
		for _, review := range reviews {
			totalRating += float64(review.Rating)
		}
		book.AverageRating = totalRating / float64(len(reviews))
		book.ReviewCount = len(reviews)
	} else {
		book.AverageRating = 0
		book.ReviewCount = 0
	}

	return r.db.Save(&book).Error
}

// GetReviewsWithUser loads review with user info
func (r *ReviewRepository) GetReviewsWithUser(review *models.Review) error {
	return r.db.Preload("User").First(review, review.ID).Error
}