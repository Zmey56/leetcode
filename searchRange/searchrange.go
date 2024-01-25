package searchRange

//Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
//
//If target is not found in the array, return [-1, -1].
//
//You must write an algorithm with O(log n) runtime complexity.

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] != target {
		return res
	}
	res[0] = left
	right = len(nums)
	for left < right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	res[1] = left - 1
	return res
}
