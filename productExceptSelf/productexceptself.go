package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	// Calculate the left products of each element
	leftProduct := 1
	for i := 0; i < n; i++ {
		output[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Calculate the right products of each element and multiply with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}
