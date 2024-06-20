package domain

import (
	"errors"
	"math/big"
)

// CalculateFibonacci returns the n-th Fibonacci number using big.Int to handle large numbers
func CalculateFibonacci(n int) (*big.Int, error) {
	if n < 1 {
		return nil, errors.New("input must be a positive integer")
	}
	if n == 1 || n == 2 {
		return big.NewInt(1), nil
	}

	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 3; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b, nil
}
