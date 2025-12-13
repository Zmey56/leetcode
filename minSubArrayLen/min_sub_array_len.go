package minSubArrayLen

//Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose
//sum is greater than or equal to target. If there is no such subarray, return 0 instead./
//
//
//Example 1:
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem constraint.
//Example 2:
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//Example 3:
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//Constraints:
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 104

func minSubArrayLen(target int, nums []int) int {
	result := 0
	sum := 0
	for i, num := range nums {
		if num >= target {
			return 1
		}
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				length := j - i + 1
				if result == 0 || length < result {
					result = length
				}
				break
			}
		}
		sum = 0
	}
	return result
}

func minSubArrayLenVer2(target int, nums []int) int {
	left := 0
	sum := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Shrink window from left while sum >= target
		for sum >= target {
			length := right - left + 1
			if result == 0 || length < result {
				result = length
			}
			sum -= nums[left]
			left++
		}
	}

	return result
}
