package checkPerfectNumber

import "math"

//A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself.
//A divisor of an integer x is an integer that can divide x evenly.
//
//Given an integer n, return true if n is a perfect number, otherwise return false.

func checkPerfectNumber(num int) bool {
	if num == 1 || num <= 0 {
		return false
	}

	sum := 1

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}

	return sum == num
}
