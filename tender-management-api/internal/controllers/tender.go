package controllers

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
    "tender-management-api/internal/models"
    "tender-management-api/internal/services"
)

// GetTenders получает список тендеров с возможностью фильтрации
func GetTenders(c echo.Context) error {
    serviceType := c.QueryParam("service_type")
    limit := c.QueryParam("limit")
    offset := c.QueryParam("offset")

    tenders, err := services.GetTenders(serviceType, limit, offset)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, tenders)
}

// CreateTender создает новый тендер
func CreateTender(c echo.Context) error {
    var tender models.Tender
    if err := c.Bind(&tender); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    createdTender, err := services.CreateTender(&tender)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, createdTender)
}

// GetUserTenders получает тендеры пользователя
func GetUserTenders(c echo.Context) error {
    username := c.QueryParam("username")
    limit := c.QueryParam("limit")
    offset := c.QueryParam("offset")

    tenders, err := services.GetUserTenders(username, limit, offset)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, tenders)
}

// GetTenderStatus получает статус тендера
func GetTenderStatus(c echo.Context) error {
    tenderID := c.Param("tenderId")
    username := c.QueryParam("username")

    status, err := services.GetTenderStatus(tenderID, username)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, status)
}

// UpdateTenderStatus изменяет статус тендера
func UpdateTenderStatus(c echo.Context) error {
    tenderID := c.Param("tenderId")
    status := c.QueryParam("status")
    username := c.QueryParam("username")

    updatedTender, err := services.UpdateTenderStatus(tenderID, status, username)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedTender)
}

// EditTender редактирует параметры тендера
func EditTender(c echo.Context) error {
    tenderID := c.Param("tenderId")
    username := c.QueryParam("username")
    var updates map[string]interface{}
    if err := c.Bind(&updates); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    updatedTender, err := services.EditTender(tenderID, username, updates)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, updatedTender)
}

// RollbackTender откатывает версию тендера
func RollbackTender(c echo.Context) error {
    tenderID := c.Param("tenderId")
    version, err := strconv.Atoi(c.Param("version"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid version format"})
    }
    username := c.QueryParam("username")

    rolledBackTender, err := services.RollbackTender(tenderID, version, username)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, rolledBackTender)
}
