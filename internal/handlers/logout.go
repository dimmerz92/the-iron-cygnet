package handlers

import (
	"net/http"
	"the-iron-cygnet/internal/auth"

	"github.com/labstack/echo/v4"
)

func Logout(ctx echo.Context) error {
	auth.RevokeSession(ctx)
	return ctx.Redirect(http.StatusSeeOther, "/")
}
