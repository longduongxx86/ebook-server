package main

import (
	"log"
	"main/internal/config"
	"main/internal/middleware"
	"main/internal/models"
	"main/internal/routes"
	"main/internal/utils"
	"main/internal/websocket"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Start WebSocket Hub
	go websocket.Manager.Start()

	// Initialize config
	db := config.InitDB()

	// Auto migrate config
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.CartItem{},
		&models.Review{},
		&models.Payment{},
		&models.Category{},
		&models.Notification{},
	)

	// Initialize MinIO (optional - will fail gracefully if not configured)
	if err := utils.InitMinIO(); err != nil {
		log.Printf("Warning: MinIO initialization failed: %v. File upload features may not work.", err)
	} else {
		log.Println("MinIO initialized successfully")
	}

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(middleware.CORS())

	// Background worker: auto cancel pending orders after 5 minutes without completed payment
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			<-ticker.C
			db := config.GetDB()
			cutoff := time.Now().Add(-5 * time.Minute).UnixMilli()
			var orders []models.Order
			if err := db.Where("status = ? AND created_at <= ?", "pending", cutoff).Find(&orders).Error; err != nil {
				continue
			}
			for _, o := range orders {
				var count int64
				if err := db.Model(&models.Payment{}).Where("order_id = ? AND status = ?", o.ID, "completed").Count(&count).Error; err != nil {
					continue
				}
				if count > 0 {
					continue
				}
				tx := db.Begin()
				ok := true
				var full models.Order
				if err := tx.Preload("Items").First(&full, o.ID).Error; err != nil {
					ok = false
				}
				if ok {
					for _, it := range full.Items {
						var book models.Book
						if err := tx.First(&book, it.BookID).Error; err != nil {
							ok = false
							break
						}
						book.Stock += it.Quantity
						if err := tx.Save(&book).Error; err != nil {
							ok = false
							break
						}
					}
				}
				if ok {
					full.Status = "cancelled"
					full.UpdatedAt = time.Now().UnixMilli()
					if err := tx.Save(&full).Error; err != nil {
						ok = false
					}
				}
				if ok {
					tx.Commit()
				} else {
					tx.Rollback()
				}
			}
		}
	}()

	routes.RegisterRoutes(r)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
