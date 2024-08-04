package auth

import "github.com/labstack/echo/v4"

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ValidateSession(ctx)
		return next(ctx)
	}
}
