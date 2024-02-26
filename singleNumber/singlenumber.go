package singleNumber

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
//You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
	result := make(map[int]int)

	for _, num := range nums {
		result[num]++
	}

	for key, value := range result {
		if value == 1 {
			return key
		}
	}

	return 0
}
