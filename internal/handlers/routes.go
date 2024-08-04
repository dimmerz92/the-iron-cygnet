package handlers

import (
	"the-iron-cygnet/internal/auth"

	"github.com/labstack/echo/v4"
)

func InitHandlerRoutes(e *echo.Echo) {
	e.Use(auth.CheckSession)
	e.Static("/static", "static")

	e.GET("/register", Register, auth.UnauthedOnly)
	e.POST("/register", Register, auth.UnauthedOnly)
}
