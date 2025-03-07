package database

import (
	"fmt"
	"os"
	"time"

	"github.com/SwanHtetAungPhyo/auth-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func InitDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort)

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			break
		}
		fmt.Printf("Connection attempt %d failed: %v\n", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := DB.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
}
