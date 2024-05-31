package maximumProduct

import "sort"

// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}
