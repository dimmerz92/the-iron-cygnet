package auth

import (
	"fmt"
	"net/http"
	"the-iron-cygnet/internal/utils"

	"github.com/labstack/echo/v4"
)

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ValidateSession(ctx)
		return next(ctx)
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role, ok := ctx.Get("Role").(string)
		if !ok || !utils.Contains(role, []string{"admin", "judge"}) {
			return ctx.String(http.StatusForbidden, fmt.Sprintf("%d FORBIDDEN", http.StatusForbidden))
		}
		return next(ctx)
	}
}

func UnauthedOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("UserID") != nil {
			return ctx.Redirect(http.StatusSeeOther, "/")
		}
		return next(ctx)
	}
}
