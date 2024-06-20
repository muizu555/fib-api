package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muizu555/fib_api/internal/router"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	router.SetupRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
