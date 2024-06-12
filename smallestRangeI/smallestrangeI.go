package smallestRangeI

//You are given an integer array nums and an integer k.
//
//In one operation, you can choose any index i where 0 <= i < nums.length and change nums[i] to nums[i] + x where x is
//an integer from the range [-k, k]. You can apply this operation at most once for each index i.
//
//The score of nums is the difference between the maximum and minimum elements in nums.
//
//Return the minimum score of nums after applying the mentioned operation at most once for each index in it.

func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}
	if maxNum-k <= minNum+k {
		return 0
	}
	return maxNum - k - (minNum + k)
}
