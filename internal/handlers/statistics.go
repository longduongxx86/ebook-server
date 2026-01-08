package handlers

import (
	"net/http"

	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	statsRepo *repositories.StatisticsRepository
}

func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{
		statsRepo: repositories.NewStatisticsRepository(),
	}
}

// Response structs (từ original)
type StatisticsResponse struct {
	Period       string                 `json:"period"`
	StartDate    int64                  `json:"start_date"`
	EndDate      int64                  `json:"end_date"`
	TotalRevenue int64                  `json:"total_revenue"`
	TotalCost    int64                  `json:"total_cost"`
	TotalProfit  int64                  `json:"total_profit"`
	OrdersCount  int64                  `json:"orders_count"`
	Details      []CategoryStatistics   `json:"details,omitempty"`
	PriceRanges  []PriceRangeStatistics `json:"price_ranges,omitempty"`
}

type CategoryStatistics struct {
	Category    string `json:"category"`
	Revenue     int64  `json:"revenue"`
	Cost        int64  `json:"cost"`
	Profit      int64  `json:"profit"`
	Quantity    int64  `json:"quantity"`
	OrdersCount int64  `json:"orders_count"`
}

type PriceRangeStatistics struct {
	PriceRange  string `json:"price_range"`
	MinPrice    int64  `json:"min_price"`
	MaxPrice    int64  `json:"max_price"`
	Revenue     int64  `json:"revenue"`
	Cost        int64  `json:"cost"`
	Profit      int64  `json:"profit"`
	Quantity    int64  `json:"quantity"`
	OrdersCount int64  `json:"orders_count"`
}

type UserStatistic struct {
	UserID      uint   `json:"user_id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	OrdersCount int64  `json:"orders_count"`
	TotalSpent  int64  `json:"total_spent"`
}

type BookStatistic struct {
	BookID   uint   `json:"book_id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Quantity int64  `json:"quantity"`
	Revenue  int64  `json:"revenue"`
	Cost     int64  `json:"cost"`
	Profit   int64  `json:"profit"`
}

// GetRevenueStatistics gets revenue statistics by period
func (h *StatisticsHandler) GetRevenueStatistics(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access statistics"})
		return
	}

	period := c.DefaultQuery("period", "day")
	dateStr := c.DefaultQuery("date", "")
	category := c.Query("category")

	// Calculate date range
	startDate, endDate, err := h.statsRepo.CalculateDateRange(period, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Get category statistics
	results, totalOrders, err := h.statsRepo.GetCategoryStatistics(startDate, endDate, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statistics: " + err.Error()})
		return
	}

	// Calculate totals
	var totalRevenue, totalCost int64
	var categoryStats []CategoryStatistics

	for _, r := range results {
		profit := r.Revenue - r.Cost
		totalRevenue += r.Revenue
		totalCost += r.Cost

		categoryStats = append(categoryStats, CategoryStatistics{
			Category:    r.Category,
			Revenue:     r.Revenue,
			Cost:        r.Cost,
			Profit:      profit,
			Quantity:    r.Quantity,
			OrdersCount: r.OrdersCount,
		})
	}

	response := StatisticsResponse{
		Period:       period,
		StartDate:    startDate,
		EndDate:      endDate,
		TotalRevenue: totalRevenue,
		TotalCost:    totalCost,
		TotalProfit:  totalRevenue - totalCost,
		OrdersCount:  totalOrders,
		Details:      categoryStats,
	}

	c.JSON(http.StatusOK, response)
}

// GetRevenueByPriceRange gets revenue statistics grouped by price ranges
func (h *StatisticsHandler) GetRevenueByPriceRange(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access statistics"})
		return
	}

	period := c.DefaultQuery("period", "day")
	dateStr := c.DefaultQuery("date", "")
	rangesStr := c.DefaultQuery("ranges", "0-50000,50000-100000,100000-200000,200000-500000,500000-999999999")

	// Calculate date range
	startDate, endDate, err := h.statsRepo.CalculateDateRange(period, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Get price range statistics
	results, err := h.statsRepo.GetPriceRangeStatistics(startDate, endDate, rangesStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statistics: " + err.Error()})
		return
	}

	// Calculate totals
	var totalRevenue, totalCost, totalOrders int64
	var priceRangeStats []PriceRangeStatistics

	for _, r := range results {
		profit := r.Revenue - r.Cost
		totalRevenue += r.Revenue
		totalCost += r.Cost
		totalOrders += r.OrdersCount

		priceRangeStats = append(priceRangeStats, PriceRangeStatistics{
			PriceRange:  r.Label,
			MinPrice:    r.MinPrice,
			MaxPrice:    r.MaxPrice,
			Revenue:     r.Revenue,
			Cost:        r.Cost,
			Profit:      profit,
			Quantity:    r.Quantity,
			OrdersCount: r.OrdersCount,
		})
	}

	response := StatisticsResponse{
		Period:       period,
		StartDate:    startDate,
		EndDate:      endDate,
		TotalRevenue: totalRevenue,
		TotalCost:    totalCost,
		TotalProfit:  totalRevenue - totalCost,
		OrdersCount:  totalOrders,
		PriceRanges:  priceRangeStats,
	}

	c.JSON(http.StatusOK, response)
}

// GetProfitStatistics gets profit/loss statistics
func (h *StatisticsHandler) GetProfitStatistics(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access statistics"})
		return
	}

	period := c.DefaultQuery("period", "day")
	dateStr := c.DefaultQuery("date", "")
	groupBy := c.Query("group_by")

	// Calculate date range
	startDate, endDate, err := h.statsRepo.CalculateDateRange(period, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Get total stats
	totalRevenue, totalCost, ordersCount, err := h.statsRepo.GetTotalStats(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statistics"})
		return
	}

	response := StatisticsResponse{
		Period:       period,
		StartDate:    startDate,
		EndDate:      endDate,
		TotalRevenue: totalRevenue,
		TotalCost:    totalCost,
		TotalProfit:  totalRevenue - totalCost,
		OrdersCount:  ordersCount,
	}

	// Add grouped details if requested
	if groupBy == "category_id" {
		results, _, err := h.statsRepo.GetCategoryStatistics(startDate, endDate, "")
		if err == nil {
			var categoryStats []CategoryStatistics
			for _, r := range results {
				categoryStats = append(categoryStats, CategoryStatistics{
					Category:    r.Category,
					Revenue:     r.Revenue,
					Cost:        r.Cost,
					Profit:      r.Revenue - r.Cost,
					Quantity:    r.Quantity,
					OrdersCount: r.OrdersCount,
				})
			}
			response.Details = categoryStats
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetUserStatistics gets statistics for users
func (h *StatisticsHandler) GetUserStatistics(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access statistics"})
		return
	}

	results, err := h.statsRepo.GetUserStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user statistics: " + err.Error()})
		return
	}

	var userStats []UserStatistic
	for _, r := range results {
		userStats = append(userStats, UserStatistic{
			UserID:      r.UserID,
			FullName:    r.FullName,
			Email:       r.Email,
			OrdersCount: r.OrdersCount,
			TotalSpent:  r.TotalSpent,
		})
	}

	c.JSON(http.StatusOK, userStats)
}

// GetBookStatistics gets statistics for books
func (h *StatisticsHandler) GetBookStatistics(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access statistics"})
		return
	}

	results, err := h.statsRepo.GetBookStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book statistics: " + err.Error()})
		return
	}

	var bookStats []BookStatistic
	for _, r := range results {
		bookStats = append(bookStats, BookStatistic{
			BookID:   r.BookID,
			Title:    r.Title,
			Category: r.Category,
			Quantity: r.Quantity,
			Revenue:  r.Revenue,
			Cost:     r.Cost,
			Profit:   r.Revenue - r.Cost,
		})
	}

	c.JSON(http.StatusOK, bookStats)
}