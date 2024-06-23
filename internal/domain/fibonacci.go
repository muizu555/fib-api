package domain

import (
	"errors"
	"math/big"
)

func CalculateFibonacci(n int) (*big.Int, error) {
	if n < 1 {
		return nil, errors.New("input must be a positive integer")
	}
	if n == 1 || n == 2 {
		return big.NewInt(1), nil
	}

	first := big.NewInt(1)
	second := big.NewInt(1)
	var result *big.Int

	for i := 3; i <= n; i++ {
		result = new(big.Int).Add(first, second)
		first, second = second, result
	}
	return result, nil
}
