package workers

import (
	"time"

	"main/internal/config"
	"main/internal/models"
)

func StartOrderAutoCancelWorker() {
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
}

