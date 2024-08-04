package auth

import (
	"fmt"
	"log"
	"net/http"
	"the-iron-cygnet/database"
	"the-iron-cygnet/database/sqlc"
	"the-iron-cygnet/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateSession(ctx echo.Context, userId string) error {
	sessionId := uuid.New().String()
	expiry := time.Now().Add(time.Hour)

	if err := database.DB.Queries.CreateSession(ctx.Request().Context(), sqlc.CreateSessionParams{
		ID:     sessionId,
		Userid: userId,
		Expiry: expiry.Unix(),
	}); err != nil {
		log.Println(err)
		return fmt.Errorf(utils.ServerError)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    sessionId,
		Expires:  expiry,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return nil
}
