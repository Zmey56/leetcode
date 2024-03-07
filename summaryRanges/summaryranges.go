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

	result := [][]string{}

	count := 0
	tmpString := []string{}
	for i := 0; i < len(nums); i++ {
		log.Println("count", count, "nums[i]", nums[i], "tmpString", tmpString, "result", result)
		if nums[i] == count {
			tmpString = append(tmpString, strconv.Itoa(nums[i]))
			count++
		} else {
			result = append(result, tmpString)
			tmpString = []string{}
			tmpString = append(tmpString, strconv.Itoa(nums[i]))
			if i == len(nums)-1 {
				result = append(result, tmpString)
			}
			count = nums[i] + 1
		}
	}

	resultString := []string{}
	for _, v := range result {
		log.Println("v", v)
		if len(v) == 1 {
			resultString = append(resultString, v[0])
		} else {
			resultString = append(resultString, v[0]+"->"+v[len(v)-1])
		}
	}

	log.Println("resultString", resultString)
	return resultString

}
