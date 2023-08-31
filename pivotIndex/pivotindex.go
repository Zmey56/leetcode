package main

func pivotIndex(nums []int) int {
	totalSum := 0
	leftSum := 0

	// Calculate the total sum of all elements in the array
	for _, num := range nums {
		totalSum += num
	}

	for i, num := range nums {
		// Check if the left sum is equal to the right sum
		if leftSum == totalSum-leftSum-num {
			return i
		}
		// Update the left sum for the next iteration
		leftSum += num
	}

	return -1
}
