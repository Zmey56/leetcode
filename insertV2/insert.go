package main

// 57. Insert Interval

// Hint
// You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi]
//  represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
//  You are also given an interval newInterval = [start, end] that represents the start and end of
//  another interval.

// Insert newInterval into intervals such that intervals is still sorted in ascending order by starti
//  and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

// Return intervals after the insertion.

// Note that you don't need to modify intervals in-place. You can make a new array and return it.
func insert(intervals [][]int, newInterval []int) [][]int {
	output := [][]int{}
	if len(intervals) < 1 {
		output = append(output, newInterval)
		return output
	}
	for _, cur := range intervals {
		if cur[1] < newInterval[0] {
			output = append(output, cur)
		} else if cur[0] > newInterval[1] {
			output = append(output, newInterval)
			newInterval = cur
		} else {
			newInterval[0] = min(newInterval[0], cur[0])
			newInterval[1] = max(newInterval[1], cur[1])
		}
	}
	output = append(output, newInterval)
	return output
}
