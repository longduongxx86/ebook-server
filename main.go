package main

import (
	"log"
	"main/internal/config"
	"main/internal/handlers"
	"main/internal/middleware"
	"main/internal/models"
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
	// Public routes

	// Khởi tạo các handler
	userHandler := handlers.NewUserHandler()
	bookHandler := handlers.NewBookHandler()
	reviewHandler := handlers.NewReviewHandler()
	cartHandler := handlers.NewCartHandler()
	orderHandler := handlers.NewOrderHandler()
	paymentHandler := handlers.NewPaymentHandler()
	dashboardHandler := handlers.NewDashboardHandler()
	statisticsHandler := handlers.NewStatisticsHandler()
	notificationHandler := handlers.NewNotificationHandler()
	categoryHandler := handlers.NewCategoryHandler()

	public := r.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
		// public.POST("/auth/google", userHandler.GoogleAuth)
		// public.GET("/verify-email", userHandler.VerifyEmail)
		public.GET("/books", bookHandler.GetBooks)
		public.GET("/books/:id", bookHandler.GetBook)
		public.GET("/books/:id/reviews", reviewHandler.GetBookReviews)
		public.GET("/categories", categoryHandler.GetListCategory)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{

		// User routes
		protected.GET("/profile", userHandler.GetProfile)
		protected.PUT("/profile", userHandler.UpdateProfile)
		protected.PUT("/change-password", userHandler.ChangePassword)

		// Upload routes (không cần handler mới)
		protected.POST("/upload/avatar", handlers.UploadAvatar)
		protected.POST("/upload/book-image", handlers.UploadBookImage)

		// Book management routes (for sellers)
		protected.POST("/books", bookHandler.CreateBook)
		protected.PUT("/books/:id", bookHandler.UpdateBook)
		protected.DELETE("/books/:id", bookHandler.DeleteBook)

		// Review routes
		protected.POST("/reviews", reviewHandler.CreateReview)
		protected.PUT("/reviews/:review_id", reviewHandler.UpdateReview)
		protected.DELETE("/reviews/:review_id", reviewHandler.DeleteReview)

		// Cart routes
		protected.POST("/cart", cartHandler.AddToCart)
		protected.GET("/cart", cartHandler.GetCart)
		protected.GET("/cart/summary", cartHandler.GetCartSummary)
		protected.PUT("/cart/items/:item_id", cartHandler.UpdateCartItem)
		protected.DELETE("/cart/items/:item_id", cartHandler.RemoveFromCart)
		protected.DELETE("/cart", cartHandler.ClearCart)

		// Order routes
		protected.POST("/orders", orderHandler.CreateOrder)
		protected.POST("/orders/from-cart", orderHandler.CreateOrderFromCart)
		protected.GET("/orders", orderHandler.GetOrders)
		protected.GET("/orders/stats", orderHandler.GetOrderStats)
		protected.GET("/orders/:id", orderHandler.GetOrder)
		protected.PUT("/orders/:id/cancel", orderHandler.CancelOrder)

		// Payment routes
		protected.POST("/orders/:id/payments", paymentHandler.CreatePayment)
		protected.GET("/orders/:id/payment", paymentHandler.GetPayment)
		protected.GET("/payments/my-payments", paymentHandler.GetUserPayments)

		// Admin/Manager routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RequireManager())
		{
			// User management (admin only)
			admin.GET("/users", userHandler.GetUsers)
			admin.PUT("/users/:id", userHandler.AdminUpdateUser)
			admin.PUT("/users/:id/reset-password", userHandler.AdminResetPassword)

			// Order management (admin only)
			admin.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

			// Payment management (admin only)
			admin.GET("/payments", paymentHandler.GetAllPayments)
			admin.PUT("/payments/:payment_id/status", paymentHandler.UpdatePaymentStatus)
		}

		// Notification routes (Authenticated users)
		protected.GET("/notifications", notificationHandler.GetNotifications)
		protected.PUT("/notifications/read", notificationHandler.MarkNotificationRead)

		// Statistics routes (Manager only)
		stats := protected.Group("/statistics")
		stats.Use(middleware.RequireManager())
		{
			stats.GET("/revenue", statisticsHandler.GetRevenueStatistics)
			stats.GET("/price-range", statisticsHandler.GetRevenueByPriceRange)
			stats.GET("/profit", statisticsHandler.GetProfitStatistics)
			stats.GET("/users", statisticsHandler.GetUserStatistics)
			stats.GET("/books", statisticsHandler.GetBookStatistics)
		}

		// Dashboard routes (Manager only)
		// Keeping them at root of /api or grouping?
		// Old paths were /api/summary, /api/charts etc.
		// Let's keep them somewhat consistent but maybe grouped is better.
		// User said "404", so reverting to old paths might be safest.
		manager := protected.Group("/")
		manager.Use(middleware.RequireManager())
		{
			manager.GET("/summary", dashboardHandler.GetDashboardSummary)
			manager.GET("/charts", dashboardHandler.GetSalesChartData)
			manager.GET("/top-selling", dashboardHandler.GetTopSelling)
			manager.GET("/recent-orders", dashboardHandler.GetRecentOrders)
			manager.GET("/metrics", dashboardHandler.GetDashboardMetrics)
		}
	}

	// WebSocket route
	r.GET("/ws", websocket.ServeWS)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
