package findDisappearedNumbers

//Given an array nums of n integers where nums[i] is in the range [1, n],
//return an array of all the integers in the range [1, n] that do not appear in nums.

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	appearNumber := make(map[int]bool)
	for _, v := range nums {
		appearNumber[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !appearNumber[i] {
			result = append(result, i)
		}
	}

	return result
}
