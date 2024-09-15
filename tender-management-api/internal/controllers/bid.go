package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "tender-management-api/internal/models"
    "tender-management-api/internal/services"
)

// GetBids получает список предложений для указанного тендера
func GetBids(c echo.Context) error {
    tenderID := c.QueryParam("tenderId")
    limit := c.QueryParam("limit")
    offset := c.QueryParam("offset")

    bids, err := services.GetBids(tenderID, limit, offset)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, bids)
}

// CreateBid создает новое предложение
func CreateBid(c echo.Context) error {
    var bid models.Bid
    if err := c.Bind(&bid); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    createdBid, err := services.CreateBid(&bid)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, createdBid)
}

// GetBid получает предложение по его ID
func GetBid(c echo.Context) error {
    bidID := c.Param("bidId")

    bid, err := services.GetBid(bidID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, bid)
}

// UpdateBid обновляет предложение
func UpdateBid(c echo.Context) error {
    bidID := c.Param("bidId")
    var updates map[string]interface{}
    if err := c.Bind(&updates); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    updatedBid, err := services.UpdateBid(bidID, updates)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedBid)
}

// DeleteBid удаляет предложение
func DeleteBid(c echo.Context) error {
    bidID := c.Param("bidId")

    if err := services.DeleteBid(bidID); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Bid deleted successfully"})
}
