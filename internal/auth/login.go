package auth

import (
	"fmt"
	"log"
	"net/mail"
	"the-iron-cygnet/database"
	"the-iron-cygnet/internal/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Email    string
	Password []byte
}

func ParseLoginForm(ctx echo.Context) (*UserLogin, error) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if len(email) < 1 {
		return nil, fmt.Errorf("Email required")
	} else if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("Not a valid email")
	}

	if len(password) < 1 {
		return nil, fmt.Errorf("Password required")
	}

	return &UserLogin{
		Email:    email,
		Password: []byte(password),
	}, nil
}

func Login(ctx echo.Context) error {
	form, err := ParseLoginForm(ctx)
	if err != nil {
		return err
	}

	user, err := database.DB.Queries.GetUserByEmail(ctx.Request().Context(), form.Email)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Incorrect email or password")
	}

	if bcrypt.CompareHashAndPassword(user.Password, form.Password) != nil {
		return fmt.Errorf("Incorrect email or password")
	}

	if err := CreateSession(ctx, user.ID); err != nil {
		return fmt.Errorf(utils.ServerError)
	}

	return nil
}
