package auth

import (
	"fmt"
	"log"
	"net/mail"
	"the-iron-cygnet/database"
	"the-iron-cygnet/database/sqlc"
	"the-iron-cygnet/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistration struct {
	FirstName, LastName, Email string
	Password                   []byte
}

func ParseRegisterForm(ctx echo.Context) (*UserRegistration, error) {
	firstName := ctx.FormValue("first-name")
	lastName := ctx.FormValue("last-name")
	email := ctx.FormValue("email")
	emailConf := ctx.FormValue("email-conf")
	password := ctx.FormValue("password")
	passwordConf := ctx.FormValue("password-conf")

	if len(firstName) < 2 || len(lastName) < 2 {
		return nil, fmt.Errorf("First and last name must be atleast 2 characters")
	}

	if len(email) < 1 || len(emailConf) < 1 {
		return nil, fmt.Errorf("Email required")
	} else if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("Not a valid email")
	} else if email != emailConf {
		return nil, fmt.Errorf("Emails do not match")
	}

	if len(password) < 8 || len(passwordConf) < 8 {
		return nil, fmt.Errorf("Password must be atleast 8 characters")
	} else if password != passwordConf {
		return nil, fmt.Errorf("Passwords do not match")
	}

	return &UserRegistration{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  []byte(password),
	}, nil
}

func Register(ctx echo.Context) error {
	form, err := ParseRegisterForm(ctx)
	if err != nil {
		return err
	}

	tx, err := database.DB.Conn.Begin()
	if err != nil {
		log.Println(err)
		return fmt.Errorf(utils.ServerError)
	}
	defer tx.Rollback()

	qtx := database.DB.Queries.WithTx(tx)

	if exists, _ := qtx.CheckEmailExistence(ctx.Request().Context(), form.Email); exists > 0 {
		return fmt.Errorf("Email already registered, try logging in")
	}

	userId := uuid.New().String()
	passwordHash, err := bcrypt.GenerateFromPassword(form.Password, 12)
	if err != nil {
		return fmt.Errorf(utils.ServerError)
	}

	if err := qtx.CreateUser(ctx.Request().Context(), sqlc.CreateUserParams{
		ID:        userId,
		Firstname: form.FirstName,
		Lastname:  form.LastName,
		Email:     form.Email,
		Password:  passwordHash,
	}); err != nil {
		return fmt.Errorf(utils.ServerError)
	}

	tx.Commit()

	return nil
}
