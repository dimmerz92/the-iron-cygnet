package handlers

import (
	"net/http"
	"the-iron-cygnet/internal/auth"
	"the-iron-cygnet/pages"

	"github.com/labstack/echo/v4"
)

func Register(ctx echo.Context) error {
	switch ctx.Request().Method {
	case http.MethodPost:
		if err := auth.Register(ctx); err != nil {
			return Render(ctx, http.StatusBadRequest, pages.Register(
				pages.RegisterProps{Error: err.Error()},
			))
		}
		return ctx.Redirect(http.StatusSeeOther, "/")

	default:
		return Render(ctx, http.StatusOK, pages.Register(pages.RegisterProps{}))
	}
}
