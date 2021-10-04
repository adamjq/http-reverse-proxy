package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/service_a/", hello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	addr := fmt.Sprintf(":%v", port)
	e.Logger.Fatal(e.Start(addr))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, from Service A!")
}
