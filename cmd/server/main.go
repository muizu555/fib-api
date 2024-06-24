package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muizu555/fib_api/internal/handler"
	"github.com/muizu555/fib_api/internal/router"
	"github.com/muizu555/fib_api/internal/usecase"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fibonacciUseCase := usecase.NewFibonacciUseCase()
	fibonacciHandler := handler.NewFibonacciHandler(fibonacciUseCase)
	// Routes
	router.SetupRouter(e, fibonacciHandler)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
