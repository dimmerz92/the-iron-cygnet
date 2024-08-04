package handlers

import (
	"net/http"
	"the-iron-cygnet/pages"
	"the-iron-cygnet/pages/layout"

	"github.com/labstack/echo/v4"
)

func Home(ctx echo.Context) error {
	role, loggedIn := ctx.Get("Role").(string)
	if !loggedIn {
		role = ""
	}

	return Render(ctx, http.StatusOK, pages.Home(pages.HomeProps{
		Layout: layout.LayoutProps{
			Nav: layout.NavProps{LoggedIn: loggedIn, Role: role},
		},
	}))
}
