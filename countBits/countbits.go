package countBits

import "log"

//Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
//ans[i] is the number of 1's in the binary representation of i.

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		log.Println(i, countOnes(i))
		result[i] = countOnes(i)
	}
	return result

}

func countOnes(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
