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

func RevokeSession(ctx echo.Context) {
	sessionCookie, err := ctx.Cookie("session")
	if err != nil {
		log.Println(err)
		return
	}

	if err := database.DB.Queries.DeleteSession(ctx.Request().Context(), sessionCookie.Value); err != nil {
		log.Println(err)
		return
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    "revoked",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

func ValidateSession(ctx echo.Context) bool {
	if err := database.DB.Queries.DeleteExpiredSessions(ctx.Request().Context()); err != nil {
		log.Println(err)
		return false
	}

	sessionCookie, err := ctx.Cookie("session")
	if err != nil {
		return false
	}

	user, err := database.DB.Queries.GetSession(ctx.Request().Context(), sessionCookie.Value)
	if err != nil {
		log.Println(err)
		return false
	} else if user.Expiry <= time.Now().Unix() {
		return false
	}

	ctx.Set("UserID", user.Userid.String)
	ctx.Set("Role", user.Role.String)

	return true
}
