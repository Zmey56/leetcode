package findLHS

//We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
//
//Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.
//
//A subsequence of array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements.

func findLHS(nums []int) int {
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}
	max := 0
	for k, v := range numsMap {
		if numsMap[k+1] != 0 && v+numsMap[k+1] > max {
			max = v + numsMap[k+1]
		}
	}
	return max
}
