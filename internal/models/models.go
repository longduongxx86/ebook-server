package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt int64          `json:"created_at" gorm:"index"`
	UpdatedAt int64          `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if b.CreatedAt == 0 {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now().UnixMilli()
	return
}

type User struct {
	BaseModel
	Email             string `json:"email" gorm:"unique;not null;index"`
	PasswordHash      string `json:"-" gorm:"not null"`
	FullName          string `json:"full_name"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	AvatarURL         string `json:"avatar_url"`
	Role              string `json:"role" gorm:"default:'customer'"`
	EmailVerified     bool   `json:"email_verified" gorm:"default:false"`
	VerificationToken string `json:"-" gorm:"index"`
}

type Book struct {
	BaseModel
	Title         string  `json:"title" gorm:"not null;index"`
	Author        string  `json:"author"`
	Description   string  `json:"description" gorm:"type:text"`
	Price         int64   `json:"price" gorm:"not null"`
	Cost          int64   `json:"cost" gorm:"default:0"`
	Stock         int     `json:"stock" gorm:"default:0"`
	Slug          string  `json:"slug" gorm:"unique;index"`
	ImageURL      string  `json:"image_url"`
	ISBN          string  `json:"isbn" gorm:"unique;index"`
	AverageRating float64 `json:"average_rating" gorm:"default:0"`
	ReviewCount   int     `json:"review_count" gorm:"default:0"`

	// Relationships (cần bổ sung từ schema)
	CategoryID *uint     `json:"category_id" gorm:"index"` // Nullable
	Category   *Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type Cart struct {
	BaseModel
	UserID uint       `json:"user_id" gorm:"not null;index"`
	User   User       `json:"user" gorm:"foreignKey:UserID"` // Thêm relationship
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	BaseModel
	CartID   uint `json:"cart_id" gorm:"not null;index"`
	BookID   uint `json:"book_id" gorm:"not null;index"`
	Quantity int  `json:"quantity" gorm:"not null"`
	Book     Book `json:"book" gorm:"foreignKey:BookID"`
	Cart     Cart `json:"-" gorm:"foreignKey:CartID"`
}

type Order struct {
	BaseModel
	OrderNumber     string      `json:"order_number" gorm:"unique;not null;index"`
	BuyerID         uint        `json:"buyer_id" gorm:"not null;index"`
	Buyer           User        `json:"buyer" gorm:"foreignKey:BuyerID"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	TotalAmount     int64       `json:"total_amount" gorm:"not null"`
	Status          string      `json:"status" gorm:"default:'pending'"`
	ShippingAddress string      `json:"shipping_address"`
	Notes           string      `json:"notes"`
}

type OrderItem struct {
	BaseModel
	OrderID   uint  `json:"order_id" gorm:"not null;index"`
	BookID    uint  `json:"book_id" gorm:"not null;index"`
	Book      Book  `json:"book" gorm:"foreignKey:BookID"`
	Quantity  int   `json:"quantity" gorm:"not null"`
	Price     int64 `json:"price" gorm:"not null"` // price snapshot (giá bán)
	Cost      int64 `json:"cost" gorm:"default:0"` // cost snapshot (vốn)
	CreatedAt int64 `json:"created_at" gorm:"index"`
	UpdatedAt int64 `json:"updated_at" gorm:"index"`
}

type Review struct {
	BaseModel
	BookID   uint   `json:"book_id" gorm:"not null;index"`
	UserID   uint   `json:"user_id" gorm:"not null;index"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Rating   int    `json:"rating" gorm:"not null"`
	Comment  string `json:"comment" gorm:"type:text"`
	Approved bool   `json:"approved" gorm:"default:true"`
}

type Payment struct {
	BaseModel
	OrderID       uint   `json:"order_id" gorm:"not null;index"`
	Order         Order  `json:"order" gorm:"foreignKey:OrderID"`
	Amount        int64  `json:"amount" gorm:"not null"`
	Method        string `json:"method"`                          // e.g., "qr","cash","bank_transfer"
	Status        string `json:"status" gorm:"default:'pending'"` // pending, completed, failed
	TransactionID string `json:"transaction_id"`
	QRCode        string `json:"qr_code" gorm:"type:text"`
	BankInfo      string `json:"bank_info" gorm:"type:text"`
	CreatedAt     int64  `json:"created_at" gorm:"index"`
	UpdatedAt     int64  `json:"updated_at" gorm:"index"`
}

type Category struct {
	BaseModel
	Name        string `json:"name" gorm:"unique;not null;index"`
	Slug        string `json:"slug" gorm:"unique;not null;index"`
	Description string `json:"description" gorm:"type:text"`
}

type Notification struct {
	BaseModel
	Title       string `json:"title" gorm:"not null"`
	Message     string `json:"message" gorm:"type:text"`
	Type        string `json:"type"` // "order", "stock", "system"
	IsRead      bool   `json:"is_read" gorm:"default:false"`
	ReferenceID uint   `json:"reference_id"` // OrderID or BookID
}
