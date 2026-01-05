package mergeV2

import "sort"

//56. Merge Intervals

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array
//of the non-overlapping intervals that cover all the intervals in the input.

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr[0] <= result[len(result)-1][1] {
			if curr[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = curr[1]
			}
		} else {
			result = append(result, curr)
		}
	}

	return result

}
