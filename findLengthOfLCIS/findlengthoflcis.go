package findLengthOfLCIS

//Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray).
//The subsequence must be strictly increasing.
//
//A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is
//[nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	countLen := 1
	prevVal := nums[0]
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > prevVal {
			countLen++
			prevVal = nums[i]
		} else {
			if countLen > maxLen {
				maxLen = countLen
			}
			countLen = 1
			prevVal = nums[i]
		}
	}

	if countLen > maxLen {
		maxLen = countLen
	}

	return maxLen
}
