package main

import (
    "log"
    "net/http"
    "tender-management-api/internal/db"
    "tender-management-api/internal/routes"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/joho/godotenv"
)

func main() {
    // Загрузка переменных окружения из .env файла
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Инициализация базы данных
    _, err := db.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Создание экземпляра Echo
    e := echo.New()

    // Настройка middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Инициализация маршрутов
    routes.InitRoutes(e)

    // Запуск сервера
    serverAddress := "0.0.0.0:8080"
    if err := e.Start(serverAddress); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Failed to start server: %v", err)
    }
}
