package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
