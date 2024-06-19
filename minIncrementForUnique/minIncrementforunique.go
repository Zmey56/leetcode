package minIncrementForUnique

import "sort"

//You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.
//
//Return the minimum number of moves to make every value in nums unique.
//
//The test cases are generated so that the answer fits in a 32-bit integer.

func minIncrementForUnique(nums []int) int {
	//sort.Ints(nums)
	valMap := make(map[int]struct{})
	count := 0

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for {
			if _, ok := valMap[tmp]; ok {
				count++
				tmp++
			} else {
				valMap[tmp] = struct{}{}
				break
			}
		}
	}
	return count

}

func minIncrementForUniqueVer2(nums []int) int {
	sort.Ints(nums)
	moves := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			increment := nums[i-1] - nums[i] + 1
			nums[i] += increment
			moves += increment
		}
	}

	return moves

}
