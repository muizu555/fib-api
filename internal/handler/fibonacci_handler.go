package handler

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/fib_api/internal/usecase"
)

type FibonacciHandler struct {
	UseCase *usecase.FibonacciUseCase
}

func NewFibonacciHandler(useCase *usecase.FibonacciUseCase) *FibonacciHandler {
	return &FibonacciHandler{UseCase: useCase}
}

func (h *FibonacciHandler) GetFibonacci(c echo.Context) error {
	nStr := c.QueryParam("n")
	if nStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "400", "message": "Parameter 'n' is required"})
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "400", "message": "Invalid parameter 'n'"})
	}

	result, err := h.UseCase.GetFibonacci(n)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": "500", "message": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]*big.Int{"result": result})
}
