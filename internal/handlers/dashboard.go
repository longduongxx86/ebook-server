package handlers

import (
	"main/internal/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardRepo *repositories.DashboardRepository
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{
		dashboardRepo: repositories.NewDashboardRepository(),
	}
}

// Response structs (từ original)
type DashboardSummaryResponse struct {
	Summary struct {
		Today struct {
			Revenue      int64 `json:"revenue"`
			Orders       int64 `json:"orders"`
			Profit       int64 `json:"profit"`
			NewCustomers int64 `json:"new_customers"`
		} `json:"today"`
		Yesterday struct {
			Revenue int64 `json:"revenue"`
			Orders  int64 `json:"orders"`
			Profit  int64 `json:"profit"`
		} `json:"yesterday"`
		ThisMonth struct {
			Revenue int64 `json:"revenue"`
			Orders  int64 `json:"orders"`
			Profit  int64 `json:"profit"`
		} `json:"this_month"`
	} `json:"summary"`
	Alerts struct {
		PendingOrders   int64          `json:"pending_orders"`
		PendingPayments int64          `json:"pending_payments"`
		LowStockBooks   []LowStockBook `json:"low_stock_books"`
	} `json:"alerts"`
	QuickStats struct {
		TotalBooks      int64 `json:"total_books"`
		TotalCustomers  int64 `json:"total_customers"`
		TotalCategories int64 `json:"total_categories"`
		ActiveOrders    int64 `json:"active_orders"`
	} `json:"quick_stats"`
}

type LowStockBook struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Stock    int    `json:"stock"`
	ImageURL string `json:"image_url,omitempty"`
}

type SalesChartDataResponse struct {
	Period string           `json:"period"`
	Data   []ChartDataPoint `json:"data"`
}

type ChartDataPoint struct {
	Date     string `json:"date"`
	Revenue  int64  `json:"revenue"`
	Orders   int64  `json:"orders"`
	Profit   int64  `json:"profit"`
	Quantity int64  `json:"quantity"`
}

type TopSellingResponse struct {
	TopBooks      []TopSellingBook     `json:"top_books"`
	TopCategories []TopSellingCategory `json:"top_categories"`
}

type TopSellingBook struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Quantity int64  `json:"quantity"`
	Revenue  int64  `json:"revenue"`
	Profit   int64  `json:"profit"`
	ImageURL string `json:"image_url,omitempty"`
}

type TopSellingCategory struct {
	Category   string  `json:"category"`
	Revenue    int64   `json:"revenue"`
	Quantity   int64   `json:"quantity"`
	Profit     int64   `json:"profit"`
	Percentage float64 `json:"percentage"`
}

// GetDashboardSummary returns comprehensive dashboard summary
func (h *DashboardHandler) GetDashboardSummary(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access dashboard"})
		return
	}

	var response DashboardSummaryResponse

	// Today stats
	todayRevenue, todayCost, todayOrders, newCustomers := h.dashboardRepo.GetTodayStats()
	response.Summary.Today.Revenue = todayRevenue
	response.Summary.Today.Orders = todayOrders
	response.Summary.Today.Profit = todayRevenue - todayCost
	response.Summary.Today.NewCustomers = newCustomers

	// Yesterday stats
	yesterdayRevenue, yesterdayCost, yesterdayOrders := h.dashboardRepo.GetYesterdayStats()
	response.Summary.Yesterday.Revenue = yesterdayRevenue
	response.Summary.Yesterday.Orders = yesterdayOrders
	response.Summary.Yesterday.Profit = yesterdayRevenue - yesterdayCost

	// Month stats
	monthRevenue, monthCost, monthOrders := h.dashboardRepo.GetMonthStats()
	response.Summary.ThisMonth.Revenue = monthRevenue
	response.Summary.ThisMonth.Orders = monthOrders
	response.Summary.ThisMonth.Profit = monthRevenue - monthCost

	// Alerts
	pendingOrders, pendingPayments := h.dashboardRepo.GetAlertStats()
	response.Alerts.PendingOrders = pendingOrders
	response.Alerts.PendingPayments = pendingPayments

	// Low stock books
	threshold, _ := strconv.Atoi(c.DefaultQuery("low_stock_threshold", "10"))
	lowStockBooks, err := h.dashboardRepo.GetLowStockBooks(threshold)
	if err == nil {
		for _, book := range lowStockBooks {
			response.Alerts.LowStockBooks = append(response.Alerts.LowStockBooks, LowStockBook{
				ID:       book.ID,
				Title:    book.Title,
				Stock:    book.Stock,
				ImageURL: book.ImageURL,
			})
		}
	}

	// Quick stats
	totalBooks, totalCustomers, totalCategories, activeOrders := h.dashboardRepo.GetQuickStats()
	response.QuickStats.TotalBooks = totalBooks
	response.QuickStats.TotalCustomers = totalCustomers
	response.QuickStats.TotalCategories = totalCategories
	response.QuickStats.ActiveOrders = activeOrders

	c.JSON(http.StatusOK, response)
}

// GetSalesChartData returns chart data for dashboard
func (h *DashboardHandler) GetSalesChartData(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access dashboard"})
		return
	}

	rangeParam := c.DefaultQuery("range", "30d")
	var days int
	switch rangeParam {
	case "7d":
		days = 7
	case "90d":
		days = 90
	default:
		days = 30
	}

	now := time.Now()
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	startDate := endDate.AddDate(0, 0, -days+1)

	// Get chart data from repository
	chartData, err := h.dashboardRepo.GetChartData(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chart data"})
		return
	}

	var response SalesChartDataResponse
	response.Period = rangeParam

	for _, data := range chartData {
		response.Data = append(response.Data, ChartDataPoint{
			Date:     data.Date,
			Revenue:  data.Revenue,
			Orders:   data.Orders,
			Profit:   data.Revenue - data.Cost,
			Quantity: data.Quantity,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetTopSelling returns top selling books and categories
func (h *DashboardHandler) GetTopSelling(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access dashboard"})
		return
	}

	period := c.DefaultQuery("period", "month")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	now := time.Now()
	var startDate, endDate time.Time

	switch period {
	case "today":
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		endDate = startDate.AddDate(0, 0, 1).Add(-time.Second)
	case "week":
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		startDate = now.AddDate(0, 0, -(weekday - 1))
		startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.Local)
		endDate = startDate.AddDate(0, 0, 7).Add(-time.Second)
	case "year":
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		endDate = startDate.AddDate(1, 0, 0).Add(-time.Second)
	default: // month
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Second)
	}

	var response TopSellingResponse

	// Top selling books
	topBooks, err := h.dashboardRepo.GetTopSellingBooks(startDate, endDate, limit)
	if err == nil {
		for _, book := range topBooks {
			response.TopBooks = append(response.TopBooks, TopSellingBook{
				ID:       book.BookID,
				Title:    book.Title,
				ImageURL: book.ImageURL,
				Quantity: book.Quantity,
				Revenue:  book.Revenue,
				Profit:   book.Revenue - book.Cost,
			})
		}
	}

	// Top selling categories
	topCategories, totalRevenue, err := h.dashboardRepo.GetTopSellingCategories(startDate, endDate, limit)
	if err == nil {
		for _, cat := range topCategories {
			percentage := 0.0
			if totalRevenue > 0 {
				percentage = float64(cat.Revenue) / float64(totalRevenue) * 100
			}

			response.TopCategories = append(response.TopCategories, TopSellingCategory{
				Category:   cat.Category,
				Revenue:    cat.Revenue,
				Quantity:   cat.Quantity,
				Profit:     cat.Revenue - cat.Cost,
				Percentage: percentage,
			})
		}
	}

	c.JSON(http.StatusOK, response)
}

// GetRecentOrders returns recent orders for dashboard
func (h *DashboardHandler) GetRecentOrders(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access dashboard"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	orders, err := h.dashboardRepo.GetRecentOrders(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recent orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetDashboardMetrics returns key metrics for dashboard widgets
func (h *DashboardHandler) GetDashboardMetrics(c *gin.Context) {
	role, exists := c.Get("user_role")
	if !exists || (role != "manager" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only managers can access dashboard"})
		return
	}

	todayRevenue, yesterdayRevenue, todayOrders, totalVisitors := h.dashboardRepo.GetMetrics()

	metrics := gin.H{
		"sales_today":         0,
		"sales_yesterday":     0,
		"growth_percentage":   0,
		"conversion_rate":     0,
		"average_order_value": 0,
	}

	metrics["sales_today"] = todayRevenue
	metrics["sales_yesterday"] = yesterdayRevenue

	// Growth percentage
	if yesterdayRevenue > 0 {
		growth := float64(todayRevenue-yesterdayRevenue) / float64(yesterdayRevenue) * 100
		metrics["growth_percentage"] = growth
	}

	// Average Order Value
	if todayOrders > 0 {
		metrics["average_order_value"] = todayRevenue / todayOrders
	}

	// Conversion rate (simplified)
	if totalVisitors > 0 {
		conversionRate := float64(todayOrders) / float64(totalVisitors) * 100
		metrics["conversion_rate"] = conversionRate
	}

	c.JSON(http.StatusOK, metrics)
}
