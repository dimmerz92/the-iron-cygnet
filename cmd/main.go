package main

import (
	"os"
	"the-iron-cygnet/database"
	"the-iron-cygnet/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	database.InitDB()
}

func main() {
	e := echo.New()

	handlers.InitHandlerRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
