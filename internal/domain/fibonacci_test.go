package domain

import (
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {
	tests := []struct {
		// case名つけるか...
		n        int
		expected string
		err      bool
	}{
		{1, "1", false},
		{2, "1", false},
		{3, "2", false},
		{10, "55", false},
		{20, "6765", false},
		{50, "12586269025", false},
		{99, "218922995834555169026", false},
		{0, "", true},
		{-1, "", true},
	}

	for _, test := range tests {
		result, err := CalculateFibonacci(test.n)
		if test.err && err == nil {
			t.Errorf("expected an error for input %d, but got nil", test.n)
		}
		if !test.err && result.String() != test.expected {
			t.Errorf("expected %s for input %d, but got %s", test.expected, test.n, result.String())
		}
	}
}
