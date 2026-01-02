package summaryRangesV2

import (
	"fmt"
)

//228. Summary Ranges

//You are given a sorted unique integer array nums.
//
//A range [a,b] is the set of all integers from a to b (inclusive).
//
//Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums
//is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
//
//Each range [a,b] in the list should be output as:
//
//"a->b" if a != b
//"a" if a == b

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}

	result := []string{}
	start := nums[0]

	for i := 1; i <= n; i++ {
		if i == n || nums[i] != nums[i-1]+1 {
			if start == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			if i < n {
				start = nums[i]
			}
		}
	}

	return result
}
