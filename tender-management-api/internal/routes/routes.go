package routes

import (
    "github.com/labstack/echo/v4"
    "tender-management-api/internal/controllers"
)

func InitRoutes(e *echo.Echo) {
    e.GET("/api/ping", controllers.CheckServer)
    e.GET("/api/tenders", controllers.GetTenders)
    e.POST("/api/tenders/new", controllers.CreateTender)
    e.GET("/api/tenders/my", controllers.GetUserTenders)
    e.GET("/api/tenders/:tenderId/status", controllers.GetTenderStatus)
    e.PUT("/api/tenders/:tenderId/status", controllers.UpdateTenderStatus)
    e.PATCH("/api/tenders/:tenderId/edit", controllers.EditTender)
    e.PUT("/api/tenders/:tenderId/rollback/:version", controllers.RollbackTender)

    // e.GET("/api/bids/my", controllers.GetUserBids)              // Получить все предложения пользователя
    e.POST("/api/bids/new", controllers.CreateBid)              // Создать новое предложение
    e.GET("/api/bids/:tenderId/list", controllers.GetBids)      // Получить все предложения по тендеру
    // e.PATCH("/api/bids/:bidId/edit", controllers.EditBid)       // Редактировать предложение
    // e.PUT("/api/bids/:bidId/rollback/:version", controllers.RollbackBid) // Откат предложения к версии
    // e.GET("/api/bids/:bidId/reviews", controllers.GetBidReviews) // Получить отзывы на предложение
}
