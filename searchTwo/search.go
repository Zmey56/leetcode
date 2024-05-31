package searchTwo

//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
//If target exists, then return its index. Otherwise, return -1.
//
//You must write an algorithm with O(log n) runtime complexity.

func search(nums []int, target int) int {
	left, mid, right := 0, (len(nums)-1)/2, len(nums)-1

	return foundValue(nums, left, mid, right, target)

}

func foundValue(nums []int, left, mid, right, target int) int {
	if nums[mid] == target {
		return mid
	}

	if target > nums[right] || target < nums[left] {
		return -1
	}

	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}

	if target < nums[mid] {
		right = mid
		mid = left + (right-left)/2
		if mid == right {
			return -1
		}
	} else {
		left = mid
		mid = left + (right-left)/2
		if mid == left {
			return -1
		}
	}

	return foundValue(nums, left, mid, right, target)
}
