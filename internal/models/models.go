package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    int64          `json:"created_at" gorm:"index"`
	UpdatedAt    int64          `json:"updated_at" gorm:"index"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	Email        string         `json:"email" gorm:"unique;not null;index"`
	PasswordHash string         `json:"-" gorm:"not null"`
	FullName     string         `json:"full_name"`
	Phone        string         `json:"phone"`
	Address      string         `json:"address"`
	AvatarURL    string         `json:"avatar_url"`
	Role         string         `json:"role" gorm:"default:'customer'"`
}

type Book struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     int64          `json:"created_at" gorm:"index"`
	UpdatedAt     int64          `json:"updated_at" gorm:"index"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	Title         string         `json:"title" gorm:"not null;index"`
	Author        string         `json:"author"`
	Description   string         `json:"description" gorm:"type:text"`
	Price         int64          `json:"price" gorm:"not null"`
	Cost          int64          `json:"cost" gorm:"default:0"`
	Stock         int            `json:"stock" gorm:"default:0"`
	ImageURL      string         `json:"image_url"`
	ISBN          string         `json:"isbn" gorm:"unique;index"`
	AverageRating float64        `json:"average_rating" gorm:"default:0"`
	ReviewCount   int            `json:"review_count" gorm:"default:0"`

	// Relationships (cần bổ sung từ schema)
	CategoryID *uint     `json:"category_id" gorm:"index"` // Nullable
	Category   *Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type Cart struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt int64          `json:"created_at" gorm:"index"`
	UpdatedAt int64          `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Items     []CartItem     `json:"items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt int64          `json:"created_at" gorm:"index"`
	UpdatedAt int64          `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	CartID    uint           `json:"cart_id" gorm:"not null;index"`
	BookID    uint           `json:"book_id" gorm:"not null;index"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Book      Book           `json:"book" gorm:"foreignKey:BookID"`
	Cart      Cart           `json:"-" gorm:"foreignKey:CartID"`
}

type Conversation struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     int64          `json:"created_at" gorm:"index"`
	UpdatedAt     int64          `json:"updated_at" gorm:"index"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	UserID        uint           `json:"user_id" gorm:"uniqueIndex;not null"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	LastMessage   string         `json:"last_message"`
	LastMessageAt int64          `json:"last_message_at"`
}

type Message struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      int64          `json:"created_at" gorm:"index"`
	UpdatedAt      int64          `json:"updated_at" gorm:"index"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
	ConversationID uint           `json:"conversation_id" gorm:"not null;index"`
	Conversation   Conversation   `json:"-" gorm:"foreignKey:ConversationID"`
	SenderID       uint           `json:"sender_id" gorm:"not null"`
	Sender         User           `json:"sender" gorm:"foreignKey:SenderID"`
	Content        string         `json:"content" gorm:"type:text;not null"`
	IsRead         bool           `json:"is_read" gorm:"default:false"`
	IsAdmin        bool           `json:"is_admin" gorm:"default:false"`
}

type Order struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       int64          `json:"created_at" gorm:"index"`
	UpdatedAt       int64          `json:"updated_at" gorm:"index"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
	OrderNumber     string         `json:"order_number" gorm:"unique;not null;index"`
	BuyerID         uint           `json:"buyer_id" gorm:"not null;index"`
	Buyer           User           `json:"buyer" gorm:"foreignKey:BuyerID"`
	Items           []OrderItem    `json:"items" gorm:"foreignKey:OrderID"`
	TotalAmount     int64          `json:"total_amount" gorm:"not null"`
	Status          string         `json:"status" gorm:"default:'pending'"`
	ShippingAddress string         `json:"shipping_address"`
	Notes           string         `json:"notes"`
}

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt int64          `json:"created_at" gorm:"index"`
	UpdatedAt int64          `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	OrderID   uint           `json:"order_id" gorm:"not null;index"`
	BookID    uint           `json:"book_id" gorm:"not null;index"`
	Book      Book           `json:"book" gorm:"foreignKey:BookID"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     int64          `json:"price" gorm:"not null"`
	Cost      int64          `json:"cost" gorm:"default:0"`
}

type Review struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt int64          `json:"created_at" gorm:"index"`
	UpdatedAt int64          `json:"updated_at" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	BookID    uint           `json:"book_id" gorm:"not null;index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Rating    int            `json:"rating" gorm:"not null"`
	Comment   string         `json:"comment" gorm:"type:text"`
	Approved  bool           `json:"approved" gorm:"default:true"`
}

type Payment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     int64          `json:"created_at" gorm:"index"`
	UpdatedAt     int64          `json:"updated_at" gorm:"index"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	OrderID       uint           `json:"order_id" gorm:"not null;index"`
	Order         Order          `json:"order" gorm:"foreignKey:OrderID"`
	Amount        int64          `json:"amount" gorm:"not null"`
	Status        string         `json:"status" gorm:"default:'pending'"`
}

type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   int64          `json:"created_at" gorm:"index"`
	UpdatedAt   int64          `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"unique;not null;index"`
	Description string         `json:"description" gorm:"type:text"`
}

type Notification struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   int64          `json:"created_at" gorm:"index"`
	UpdatedAt   int64          `json:"updated_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"title" gorm:"not null"`
	Message     string         `json:"message" gorm:"type:text"`
	Type        string         `json:"type"`
	IsRead      bool           `json:"is_read" gorm:"default:false"`
	ReferenceID uint           `json:"reference_id"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if u.CreatedAt == 0 {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UnixMilli()
	return
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if b.CreatedAt == 0 {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
	return
}

func (b *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now().UnixMilli()
	return
}

func (c *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if c.CreatedAt == 0 {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
	return
}

func (c *Cart) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UnixMilli()
	return
}

func (ci *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if ci.CreatedAt == 0 {
		ci.CreatedAt = now
	}
	ci.UpdatedAt = now
	return
}

func (ci *CartItem) BeforeUpdate(tx *gorm.DB) (err error) {
	ci.UpdatedAt = time.Now().UnixMilli()
	return
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if o.CreatedAt == 0 {
		o.CreatedAt = now
	}
	o.UpdatedAt = now
	return
}

func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now().UnixMilli()
	return
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if oi.CreatedAt == 0 {
		oi.CreatedAt = now
	}
	oi.UpdatedAt = now
	return
}

func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	oi.UpdatedAt = time.Now().UnixMilli()
	return
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if r.CreatedAt == 0 {
		r.CreatedAt = now
	}
	r.UpdatedAt = now
	return
}

func (r *Review) BeforeUpdate(tx *gorm.DB) (err error) {
	r.UpdatedAt = time.Now().UnixMilli()
	return
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if p.CreatedAt == 0 {
		p.CreatedAt = now
	}
	p.UpdatedAt = now
	return
}

func (p *Payment) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now().UnixMilli()
	return
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if c.CreatedAt == 0 {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().UnixMilli()
	return
}

func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UnixMilli()
	if n.CreatedAt == 0 {
		n.CreatedAt = now
	}
	n.UpdatedAt = now
	return
}

func (n *Notification) BeforeUpdate(tx *gorm.DB) (err error) {
	n.UpdatedAt = time.Now().UnixMilli()
	return
}
