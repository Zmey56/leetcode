package thirdMax

// Given an integer array nums, return the third distinct maximum number in this array.
//If the third maximum does not exist, return the maximum number.

func thirdMax(nums []int) int {
	values := make(map[int]bool)
	count := 0
	for _, v := range nums {
		if !values[v] {
			values[v] = true
			count++
		}

		if len(values) > 2 {
			break
		}
	}

	result := 0
	for _, v := range values {
		if v
	}


}
