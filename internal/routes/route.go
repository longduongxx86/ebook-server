package routes

import (
	"main/internal/handlers"
	"main/internal/middleware"
	"main/internal/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
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
	chatHandler := handlers.NewChatHandler()

	public := r.Group("/api")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
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
		protected.POST("/logout", userHandler.Logout)

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
		protected.POST("/cart/add", cartHandler.AddToCart)
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

		// Chat routes
		protected.GET("/chat/history", chatHandler.GetChatHistory)

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

			// Chat management
			admin.GET("/chat/conversations", chatHandler.GetConversations)
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
			manager.GET("/users", userHandler.GetUsers)
			manager.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
			manager.GET("/payments/all", paymentHandler.GetAllPayments)
			manager.PUT("/payments/:payment_id/status", paymentHandler.UpdatePaymentStatus)
		}
	}

	// WebSocket route
	r.GET("/ws", websocket.ServeWS)
}
