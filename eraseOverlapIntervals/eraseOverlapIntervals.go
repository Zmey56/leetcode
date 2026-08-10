package main

import (
	"slices"
)

// 435. Non-overlapping Intervals

// Given an array of intervals intervals where intervals[i] = [starti, endi],
//  return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

// Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3]
//
//	are non-overlapping.
func eraseOverlapIntervals(intervals [][]int) int {

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[1] - b[1]
	})

	count := 0

	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			count++
		} else {
			prevEnd = intervals[i][1]
		}
	}
	return count
}
