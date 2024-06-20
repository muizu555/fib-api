package usecase

import (
	"math/big"

	"github.com/muizu555/fib_api/internal/domain"
)

// FibonacciUseCase provides methods for calculating Fibonacci numbers
type FibonacciUseCase struct{}

// NewFibonacciUseCase creates a new FibonacciUseCase
func NewFibonacciUseCase() *FibonacciUseCase {
	return &FibonacciUseCase{}
}

// GetFibonacci returns the n-th Fibonacci number
func (u *FibonacciUseCase) GetFibonacci(n int) (*big.Int, error) {
	return domain.CalculateFibonacci(n)
}
