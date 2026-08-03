package main

//912. Sort an Array

// Given an array of integers nums, sort the array in ascending order and return it.

// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity
//
//	and with the smallest space complexity possible.
func sortArray(nums []int) []int {
	// 1. Find the smallest and largest numbers to know our range
	minVal, maxVal := nums[0], nums[0]
	for _, n := range nums {
		if n < minVal {
			minVal = n
		}
		if n > maxVal {
			maxVal = n
		}
	}

	// 2. Create a "tally sheet" (array) to count occurrences
	counts := make([]int, maxVal-minVal+1)
	for _, n := range nums {
		counts[n-minVal]++ // Tally up each number
	}

	// 3. Overwrite the original array with the tallied numbers
	index := 0
	for i, count := range counts {
		for j := 0; j < count; j++ {
			nums[index] = i + minVal
			index++
		}
	}

	return nums
}
