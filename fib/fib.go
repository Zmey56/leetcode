package fib

//The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence
//, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
//
//F(0) = 0, F(1) = 1
//F(n) = F(n - 1) + F(n - 2), for n > 1.
//Given n, calculate F(n).

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	result := 0
	prevResult := 1
	for i := 2; i <= n; i++ {
		if i == 2 {
			result = 1
			continue
		}
		result, prevResult = result+prevResult, result
	}

	return result
}
