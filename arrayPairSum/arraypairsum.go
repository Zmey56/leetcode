package arrayPairSum

//Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn)
//such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.

func arrayPairSum(nums []int) int {
	// Sort the array
	nums = mergeSort(nums)

	// Sum the minimum of every pair
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum

}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
