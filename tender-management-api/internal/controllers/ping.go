package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// CheckServer обрабатывает запросы на эндпоинт /ping
func CheckServer(c echo.Context) error {
    return c.String(http.StatusOK, "ok")
}
