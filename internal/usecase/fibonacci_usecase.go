package usecase

import (
	"math/big"

	"github.com/muizu555/fib_api/internal/domain"
)

type FibonacciUseCase struct{}

func NewFibonacciUseCase() *FibonacciUseCase {
	return &FibonacciUseCase{}
}

func (u *FibonacciUseCase) GetFibonacci(n int) (*big.Int, error) {
	return domain.CalculateFibonacci(n)
}
