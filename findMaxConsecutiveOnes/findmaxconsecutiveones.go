package findMaxConsecutiveOnes

import "math"

//Given a binary array nums, return the maximum number of consecutive 1's in the array.

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			max = int(math.Max(float64(max), float64(count)))
			count = 0
		}
	}

	return int(math.Max(float64(max), float64(count)))
}
