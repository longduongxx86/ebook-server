package main

import (
	"log"
	"main/internal/config"
	"main/internal/middleware"
	"main/internal/models"
	"main/internal/routes"
	"main/internal/utils"
	"main/internal/websocket"
	"main/internal/workers"

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

	workers.StartOrderAutoCancelWorker()

	routes.RegisterRoutes(r)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
