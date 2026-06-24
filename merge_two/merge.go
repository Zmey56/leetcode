package mergetwo

import (
	"sort"
)

// 56. Merge Intervals

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
//
//	and return an array of the non-overlapping intervals that cover all the intervals in the input.
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastIndex := 0

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[lastIndex][1] {
			if intervals[i][1] > intervals[lastIndex][1] {
				intervals[lastIndex][1] = intervals[i][1]
			}
		} else {
			lastIndex++
			intervals[lastIndex] = intervals[i]
		}
	}

	return intervals[:lastIndex+1]
}
