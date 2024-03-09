package summaryRanges

import (
	"log"
	"strconv"
)

// You are given a sorted unique integer array nums.
//
//A range [a,b] is the set of all integers from a to b (inclusive).
//
//Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
//That is, each element of nums is covered by exactly one of the ranges,
//and there is no integer x such that x is in one of the ranges but not in nums.
//
//Each range [a,b] in the list should be output as:
//
//"a->b" if a != b
//"a" if a == b

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var result []string
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			i++
		}
		if start != nums[i] {
			result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		} else {
			result = append(result, strconv.Itoa(start))
		}
	}
	log.Println(result)
	return result

}
