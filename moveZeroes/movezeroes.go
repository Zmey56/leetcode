package main

func moveZeroes(nums []int) {
	result := []int{}
	tmpZero := []int{}

	for _, i := range nums {
		if i != 0 {
			result = append(result, i)
		} else {
			tmpZero = append(tmpZero, i)
		}
	}

	result = append(result, tmpZero...)

	copy(nums, result)
}
