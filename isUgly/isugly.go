package isUgly

import "log"

//An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
//
//Given an integer n, return true if n is an ugly number.

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	log.Println("N", n)
	for _, i := range []int{2, 3, 5} {
		for n%i == 0 {
			log.Println(n, i)
			n /= i
		}
	}

	return n == 1

}
