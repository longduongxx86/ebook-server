package repositories

import (
	"main/internal/config"
	"main/internal/models"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type StatisticsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepository() *StatisticsRepository {
	return &StatisticsRepository{
		db: config.GetDB(),
	}
}

// Helper: Calculate date range based on period
func (r *StatisticsRepository) CalculateDateRange(period, dateStr string) (int64, int64, error) {
	now := time.Now()
	var startDate, endDate int64

	if dateStr == "" {
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		startDate = today.UnixMilli()
		endDate = today.AddDate(0, 0, 1).UnixMilli() - 1
	} else {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return 0, 0, err
		}
		startDate = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
		endDate = startDate + 24*60*60*1000 - 1
	}

	switch period {
	case "week":
		weekStart := time.Unix(startDate/1000, 0)
		weekday := int(weekStart.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		weekStart = weekStart.AddDate(0, 0, -(weekday - 1))
		startDate = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
		endDate = startDate + 7*24*60*60*1000 - 1
	case "month":
		monthStart := time.Unix(startDate/1000, 0)
		startDate = time.Date(monthStart.Year(), monthStart.Month(), 1, 0, 0, 0, 0, time.Local).UnixMilli()
		nextMonth := time.Date(monthStart.Year(), monthStart.Month()+1, 1, 0, 0, 0, 0, time.Local)
		endDate = nextMonth.UnixMilli() - 1
	}

	return startDate, endDate, nil
}

// GetCategoryStatistics gets statistics by category
func (r *StatisticsRepository) GetCategoryStatistics(startDate, endDate int64, category string) ([]struct {
	Category    string
	Revenue     int64
	Cost        int64
	Quantity    int64
	OrdersCount int64
}, int64, error) {
	query := r.db.Table("order_items").
		Select(`
			COALESCE(categories.name, 'Uncategorized') as category,
			SUM(order_items.price * order_items.quantity) as revenue,
			SUM(order_items.cost * order_items.quantity) as cost,
			SUM(order_items.quantity) as quantity,
			COUNT(DISTINCT order_items.order_id) as orders_count
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Joins("JOIN books ON books.id = order_items.book_id").
		Joins("LEFT JOIN categories ON categories.id = books.category_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?", startDate, endDate).
		Group("COALESCE(categories.name, 'Uncategorized')")

	if category != "" {
		query = query.Where("categories.name = ?", category)
	}

	var results []struct {
		Category    string
		Revenue     int64
		Cost        int64
		Quantity    int64
		OrdersCount int64
	}

	err := query.Scan(&results).Error

	// Get total orders count - SỬA DÒNG NÀY
	var totalOrders int64
	r.db.Model(&models.Order{}).Where("status IN (?) AND created_at >= ? AND created_at <= ?",
		[]string{"confirmed", "paid", "shipped", "delivered"}, startDate, endDate).
		Count(&totalOrders)

	return results, totalOrders, err
}

// GetPriceRangeStatistics gets statistics by price range
func (r *StatisticsRepository) GetPriceRangeStatistics(startDate, endDate int64, rangesStr string) ([]struct {
	MinPrice    int64
	MaxPrice    int64
	Label       string
	Revenue     int64
	Cost        int64
	Quantity    int64
	OrdersCount int64
}, error) {
	ranges := r.parsePriceRanges(rangesStr)

	var results []struct {
		MinPrice    int64
		MaxPrice    int64
		Label       string
		Revenue     int64
		Cost        int64
		Quantity    int64
		OrdersCount int64
	}

	for _, priceRange := range ranges {
		var result struct {
			Revenue     int64
			Cost        int64
			Quantity    int64
			OrdersCount int64
		}

		err := r.db.Table("order_items").
			Select(`
				SUM(order_items.price * order_items.quantity) as revenue,
				SUM(order_items.cost * order_items.quantity) as cost,
				SUM(order_items.quantity) as quantity,
				COUNT(DISTINCT order_items.order_id) as orders_count
			`).
			Joins("JOIN orders ON orders.id = order_items.order_id").
			Joins("JOIN books ON books.id = order_items.book_id").
			Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
			Where("orders.created_at >= ? AND orders.created_at <= ?", startDate, endDate).
			Where("order_items.price >= ? AND order_items.price < ?", priceRange.MinPrice, priceRange.MaxPrice).
			Scan(&result).Error

		if err != nil {
			return nil, err
		}

		results = append(results, struct {
			MinPrice    int64
			MaxPrice    int64
			Label       string
			Revenue     int64
			Cost        int64
			Quantity    int64
			OrdersCount int64
		}{
			MinPrice:    priceRange.MinPrice,
			MaxPrice:    priceRange.MaxPrice,
			Label:       priceRange.Label,
			Revenue:     result.Revenue,
			Cost:        result.Cost,
			Quantity:    result.Quantity,
			OrdersCount: result.OrdersCount,
		})
	}

	return results, nil
}

// GetTotalStats gets total revenue, cost, and profit
func (r *StatisticsRepository) GetTotalStats(startDate, endDate int64) (revenue, cost, ordersCount int64, err error) {
	var totalStats struct {
		Revenue int64
		Cost    int64
	}

	err = r.db.Table("order_items").
		Select(`
			SUM(order_items.price * order_items.quantity) as revenue,
			SUM(order_items.cost * order_items.quantity) as cost
		`).
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Where("orders.created_at >= ? AND orders.created_at <= ?", startDate, endDate).
		Scan(&totalStats).Error

	if err != nil {
		return 0, 0, 0, err
	}

	// SỬA DÒNG NÀY
	err = r.db.Model(&models.Order{}).Where("status IN (?) AND created_at >= ? AND created_at <= ?",
		[]string{"confirmed", "paid", "shipped", "delivered"}, startDate, endDate).
		Count(&ordersCount).Error

	return totalStats.Revenue, totalStats.Cost, ordersCount, err
}

// GetUserStatistics gets user statistics
func (r *StatisticsRepository) GetUserStatistics() ([]struct {
	UserID      uint
	FullName    string
	Email       string
	OrdersCount int64
	TotalSpent  int64
}, error) {
	var results []struct {
		UserID      uint
		FullName    string
		Email       string
		OrdersCount int64
		TotalSpent  int64
	}

	err := r.db.Table("users").
		Select(`
			users.id as user_id,
			users.full_name,
			users.email,
			COUNT(DISTINCT orders.id) as orders_count,
			COALESCE(SUM(orders.total_amount), 0) as total_spent
		`).
		Joins("LEFT JOIN orders ON orders.buyer_id = users.id AND orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Group("users.id").
		Order("total_spent DESC").
		Scan(&results).Error

	return results, err
}

// GetBookStatistics gets book statistics
func (r *StatisticsRepository) GetBookStatistics() ([]struct {
	BookID   uint
	Title    string
	Category string
	Quantity int64
	Revenue  int64
	Cost     int64
}, error) {
	var results []struct {
		BookID   uint
		Title    string
		Category string
		Quantity int64
		Revenue  int64
		Cost     int64
	}

	err := r.db.Table("books").
		Select(`
			books.id as book_id,
			books.title,
			COALESCE(categories.name, 'Uncategorized') as category,
			COALESCE(SUM(order_items.quantity), 0) as quantity,
			COALESCE(SUM(order_items.price * order_items.quantity), 0) as revenue,
			COALESCE(SUM(order_items.cost * order_items.quantity), 0) as cost
		`).
		Joins("LEFT JOIN order_items ON order_items.book_id = books.id").
		Joins("LEFT JOIN orders ON orders.id = order_items.order_id AND orders.status IN (?)", []string{"confirmed", "paid", "shipped", "delivered"}).
		Joins("LEFT JOIN categories ON categories.id = books.category_id").
		Group("books.id").
		Order("revenue DESC").
		Scan(&results).Error

	return results, err
}

// Helper: Parse price ranges
func (r *StatisticsRepository) parsePriceRanges(rangesStr string) []struct {
	MinPrice int64
	MaxPrice int64
	Label    string
} {
	ranges := []struct {
		MinPrice int64
		MaxPrice int64
		Label    string
	}{}
	parts := strings.Split(rangesStr, ",")

	for _, part := range parts {
		pair := strings.Split(part, "-")
		if len(pair) == 2 {
			min, err1 := strconv.ParseInt(pair[0], 10, 64)
			max, err2 := strconv.ParseInt(pair[1], 10, 64)
			if err1 == nil && err2 == nil {
				label := r.formatPriceRange(min, max)
				ranges = append(ranges, struct {
					MinPrice int64
					MaxPrice int64
					Label    string
				}{
					MinPrice: min,
					MaxPrice: max,
					Label:    label,
				})
			}
		}
	}

	if len(ranges) == 0 {
		ranges = []struct {
			MinPrice int64
			MaxPrice int64
			Label    string
		}{
			{MinPrice: 0, MaxPrice: 50000, Label: "0-50,000"},
			{MinPrice: 50000, MaxPrice: 100000, Label: "50,000-100,000"},
			{MinPrice: 100000, MaxPrice: 200000, Label: "100,000-200,000"},
			{MinPrice: 200000, MaxPrice: 500000, Label: "200,000-500,000"},
			{MinPrice: 500000, MaxPrice: 999999999, Label: "500,000+"},
		}
	}

	return ranges
}

// Helper: Format price range label
func (r *StatisticsRepository) formatPriceRange(min, max int64) string {
	if max >= 999999999 {
		return r.formatPrice(min) + "+"
	}
	return r.formatPrice(min) + "-" + r.formatPrice(max)
}

// Helper: Format price with commas
func (r *StatisticsRepository) formatPrice(price int64) string {
	s := strconv.FormatInt(price, 10)
	n := len(s)
	if n <= 3 {
		return s
	}
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(s[i])
	}
	return result
}
