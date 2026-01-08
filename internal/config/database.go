package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes database connection
func InitDB() *gorm.DB {
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbUser := GetEnv("DB_USER", "root")
	dbPassword := GetEnv("DB_PASSWORD", "rootpassword")
	dbName := GetEnv("DB_NAME", "ebook_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	return db
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
