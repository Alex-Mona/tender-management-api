package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "time"
	"fmt"
    "tender-management-api/internal/models"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB
var ErrRecordNotFound = gorm.ErrRecordNotFound

// InitDB инициализирует подключение к базе данных и выполняет миграции.
func InitDB() (*gorm.DB, error) {
    dsn := os.Getenv("POSTGRES_CONN")
    if dsn == "" {
        return nil, fmt.Errorf("POSTGRES_CONN environment variable is not set")
    }

    // Проверка строки подключения
    fmt.Println("Connection string:", dsn)

    // Создание подключения к базе данных с логированием
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }

    // Автоматическая миграция структур
    if err := db.AutoMigrate(&models.Tender{}, &models.Bid{}, &models.BidReview{}); err != nil {
        return nil, err
    }

    // Настройка соединений
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetConnMaxIdleTime(30 * time.Second)
    sqlDB.SetConnMaxLifetime(5 * time.Minute)
    sqlDB.SetMaxOpenConns(10)
    sqlDB.SetMaxIdleConns(5)

    DB = db
    return db, nil
}
