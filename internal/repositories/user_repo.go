package repositories

import (
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: config.GetDB(),
	}
}

// FindByID finds user by ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmailExcludingID finds user by email excluding specific ID
func (r *UserRepository) FindByEmailExcludingID(email string, excludeID uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND id != ?", email, excludeID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user
func (r *UserRepository) Update(user *models.User) error {
	user.UpdatedAt = time.Now().UnixMilli()
	return r.db.Save(user).Error
}

// FindAll gets all users
func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}