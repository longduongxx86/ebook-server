package repositories

import (
	"main/internal/config"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository() *DashboardRepository {
	return &DashboardRepository{
		db: config.GetDB(),
	}
}

// GetTodayStats gets today's stats
func (r *DashboardRepository) GetTodayStats() (revenue, cost, orders, newCustomers int64) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	todayEnd := todayStart.AddDate(0, 0, 1).Add(-time.Second)

	// Today's sales stats
	var todayStats struct {
		Revenue int64
		Cost    int64
		Orders  int64
	}
	r.db.Table("order_items").
		Select(`
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
			COALESCE(SUM(order_items.cost * order_items.quantity), 0) as cost,
			COUNT(DISTINCT order_items.order_id) as orders
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			todayStart.UnixMilli(), todayEnd.UnixMilli()).
		Scan(&todayStats)

	// New customers today
	r.db.Model(&models.User{}).
		Where("created_at >= ? AND created_at <= ?",
			todayStart.UnixMilli(), todayEnd.UnixMilli()).
		Count(&newCustomers)

	return todayStats.Revenue, todayStats.Cost, todayStats.Orders, newCustomers
}

// GetYesterdayStats gets yesterday's stats
func (r *DashboardRepository) GetYesterdayStats() (revenue, cost, orders int64) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yesterdayStart := todayStart.AddDate(0, 0, -1)
	yesterdayEnd := todayStart.Add(-time.Second)

	var yesterdayStats struct {
		Revenue int64
		Cost    int64
		Orders  int64
	}
	r.db.Table("order_items").
		Select(`
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
			COALESCE(SUM(order_items.cost * order_items.quantity), 0) as cost,
			COUNT(DISTINCT order_items.order_id) as orders
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			yesterdayStart.UnixMilli(), yesterdayEnd.UnixMilli()).
		Scan(&yesterdayStats)

	return yesterdayStats.Revenue, yesterdayStats.Cost, yesterdayStats.Orders
}

// GetMonthStats gets this month's stats
func (r *DashboardRepository) GetMonthStats() (revenue, cost, orders int64) {
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	nextMonth := monthStart.AddDate(0, 1, 0)
	monthEnd := nextMonth.Add(-time.Second)

	var monthStats struct {
		Revenue int64
		Cost    int64
		Orders  int64
	}
	r.db.Table("order_items").
		Select(`
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
			COALESCE(SUM(order_items.cost * order_items.quantity), 0) as cost,
			COUNT(DISTINCT order_items.order_id) as orders
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			monthStart.UnixMilli(), monthEnd.UnixMilli()).
		Scan(&monthStats)

	return monthStats.Revenue, monthStats.Cost, monthStats.Orders
}

// GetAlertStats gets alert stats
func (r *DashboardRepository) GetAlertStats() (pendingOrders, pendingPayments int64) {
	r.db.Model(&models.Order{}).
		Where("status = ?", "pending").
		Count(&pendingOrders)

	r.db.Model(&models.Payment{}).
		Where("status = ?", "pending").
		Count(&pendingPayments)

	return pendingOrders, pendingPayments
}

// GetLowStockBooks gets low stock books
func (r *DashboardRepository) GetLowStockBooks(threshold int) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("stock <= ?", threshold).
		Order("stock ASC").
		Limit(10).
		Find(&books).Error
	return books, err
}

// GetQuickStats gets quick stats
func (r *DashboardRepository) GetQuickStats() (totalBooks, totalCustomers, totalCategories, activeOrders int64) {
	r.db.Model(&models.Book{}).Count(&totalBooks)
	r.db.Model(&models.User{}).Where("role = ?", "customer").Count(&totalCustomers)
	r.db.Model(&models.Category{}).Count(&totalCategories)
	r.db.Model(&models.Order{}).
		Where("status IN (?)", []string{"pending", "confirmed", "shipped"}).
		Count(&activeOrders)

	return totalBooks, totalCustomers, totalCategories, activeOrders
}

// GetChartData gets chart data for specified date range
func (r *DashboardRepository) GetChartData(startDate, endDate time.Time) ([]struct {
	Date     string
	Revenue  int64
	Cost     int64
	Orders   int64
	Quantity int64
}, error) {
	var result []struct {
		Date     string
	Revenue  int64
	Cost     int64
	Orders   int64
	Quantity int64
	}

	// Generate all dates in range
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		dayStart := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
		dayEnd := dayStart.AddDate(0, 0, 1).Add(-time.Second)

		var stats struct {
			Revenue  int64
			Cost     int64
			Orders   int64
			Quantity int64
		}

		r.db.Table("order_items").
			Select(`
				COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
				COALESCE(SUM(order_items.cost * order_items.quantity), 0) as cost,
				COUNT(DISTINCT order_items.order_id) as orders,
				COALESCE(SUM(order_items.quantity), 0) as quantity
			`).
			Joins("JOIN orders ON orders.id = order_items.order_id").
			Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
			Where("orders.created_at >= ? AND orders.created_at <= ?",
				dayStart.UnixMilli(), dayEnd.UnixMilli()).
			Scan(&stats)

		result = append(result, struct {
			Date     string
			Revenue  int64
			Cost     int64
			Orders   int64
			Quantity int64
		}{
			Date:     d.Format("2006-01-02"),
			Revenue:  stats.Revenue,
			Cost:     stats.Cost,
			Orders:   stats.Orders,
			Quantity: stats.Quantity,
		})
	}

	return result, nil
}

// GetTopSellingBooks gets top selling books for period
func (r *DashboardRepository) GetTopSellingBooks(startDate, endDate time.Time, limit int) ([]struct {
	BookID   uint
	Title    string
	ImageURL string
	Quantity int64
	Revenue  int64
	Cost     int64
}, error) {
	var result []struct {
		BookID   uint
		Title    string
		ImageURL string
		Quantity int64
		Revenue  int64
		Cost     int64
	}

	err := r.db.Table("order_items").
		Select(`
			books.id as book_id,
			books.title as title,
			books.image_url as image_url,
			SUM(order_items.quantity) as quantity,
			SUM(order_items.price * order_items.quantity) as revenue,
			SUM(order_items.cost * order_items.quantity) as cost
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Joins("JOIN books ON books.id = order_items.book_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			startDate.UnixMilli(), endDate.UnixMilli()).
		Group("books.id, books.title, books.image_url").
		Order("quantity DESC").
		Limit(limit).
		Scan(&result).Error

	return result, err
}

// GetTopSellingCategories gets top selling categories for period
func (r *DashboardRepository) GetTopSellingCategories(startDate, endDate time.Time, limit int) ([]struct {
	Category string
	Revenue  int64
	Quantity int64
	Cost     int64
}, int64, error) {
	var result []struct {
		Category string
		Revenue  int64
		Quantity int64
		Cost     int64
	}

	err := r.db.Table("order_items").
		Select(`
			COALESCE(categories.name, 'Uncategorized') as category,
			SUM(order_items.price * order_items.quantity) as revenue,
			SUM(order_items.quantity) as quantity,
			SUM(order_items.cost * order_items.quantity) as cost
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Joins("JOIN books ON books.id = order_items.book_id").
		Joins("LEFT JOIN categories ON categories.id = books.category_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			startDate.UnixMilli(), endDate.UnixMilli()).
		Group("COALESCE(categories.name, 'Uncategorized')").
		Order("revenue DESC").
		Limit(limit).
		Scan(&result).Error

	// Calculate total revenue
	var totalRevenue int64
	for _, cat := range result {
		totalRevenue += cat.Revenue
	}

	return result, totalRevenue, err
}

// GetRecentOrders gets recent orders
func (r *DashboardRepository) GetRecentOrders(limit int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items.Book").Preload("Buyer").
		Order("created_at DESC").
		Limit(limit).
		Find(&orders).Error
	return orders, err
}

// GetMetrics gets dashboard metrics
func (r *DashboardRepository) GetMetrics() (todayRevenue, yesterdayRevenue, todayOrders int64, totalVisitors int64) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	todayEnd := todayStart.AddDate(0, 0, 1).Add(-time.Second)
	yesterdayStart := todayStart.AddDate(0, 0, -1)
	yesterdayEnd := todayStart.Add(-time.Second)

	// Today's sales
	var todayStats struct {
		Revenue int64
		Orders  int64
	}
	r.db.Table("order_items").
		Select(`
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
			COUNT(DISTINCT order_items.order_id) as orders
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			todayStart.UnixMilli(), todayEnd.UnixMilli()).
		Scan(&todayStats)

	// Yesterday's sales
	var yesterdayStats struct {
		Revenue int64
	}
	r.db.Table("order_items").
		Select(`
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?",
			yesterdayStart.UnixMilli(), yesterdayEnd.UnixMilli()).
		Scan(&yesterdayStats)

	// Total visitors
	r.db.Model(&models.User{}).Where("role = ?", "customer").Count(&totalVisitors)

	return todayStats.Revenue, yesterdayStats.Revenue, todayStats.Orders, totalVisitors
}