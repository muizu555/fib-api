package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muizu555/fib_api/internal/handler"
)

func SetupRouter(e *echo.Echo, fibonacciHandler *handler.FibonacciHandler) {

	e.GET("/fib", fibonacciHandler.GetFibonacci)
}
