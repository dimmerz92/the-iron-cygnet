package handlers

import (
	"net/http"
	"the-iron-cygnet/internal/auth"
	"the-iron-cygnet/pages"

	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	switch ctx.Request().Method {
	case http.MethodPost:
		if err := auth.Login(ctx); err != nil {
			return Render(ctx, http.StatusBadRequest, pages.Login(pages.LoginProps{
				Error: err.Error(),
			}))
		}
		return ctx.Redirect(http.StatusSeeOther, "/")

	default:
		return Render(ctx, http.StatusOK, pages.Login(pages.LoginProps{}))
	}
}
