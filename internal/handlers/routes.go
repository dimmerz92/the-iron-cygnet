package handlers

import "github.com/labstack/echo/v4"

func InitHandlerRoutes(e *echo.Echo) {
	e.Static("/static", "static")
}
