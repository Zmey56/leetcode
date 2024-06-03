package isMonotonic

// An array is monotonic if it is either monotone increasing or monotone decreasing.
//
//An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if
//for all i <= j, nums[i] >= nums[j].
//
//Given an integer array nums, return true if the given array is monotonic, or false otherwise.

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	prev := nums[0]

	// check if the array is increasing
	increasing := true
	for i := 1; i < len(nums); i++ {
		if nums[i] < prev {
			increasing = false
			break
		}
		prev = nums[i]
	}

	// check if the array is decreasing

	decreasing := true
	prev = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > prev {
			decreasing = false
			break
		}
		prev = nums[i]
	}

	return increasing || decreasing
}
