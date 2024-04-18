package fib

import "testing"

// Test for function fib
func Test_fib(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}

	for _, tt := range tests {
		if got := fib(tt.n); got != tt.result {
			t.Errorf("fib(%v) = %v, want %v", tt.n, got, tt.result)
		}
	}
}
