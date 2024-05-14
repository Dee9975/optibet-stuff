package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	app := echo.New()

	app.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
