package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muizu555/fib_api/internal/handler"
	"github.com/muizu555/fib_api/internal/usecase"
)

func SetupRouter(e *echo.Echo) {
	fibonacciUseCase := usecase.NewFibonacciUseCase()
	//　依存性注入
	fibonacciHandler := handler.NewFibonacciHandler(fibonacciUseCase)

	e.GET("/fib", fibonacciHandler.GetFibonacci)
}
