package insert

//57. Insert Interval

//You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start
//and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval
//newInterval = [start, end] that represents the start and end of another interval.
//
//Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still
//does not have any overlapping intervals (merge overlapping intervals if necessary).
//
//Return intervals after the insertion.
//
//Note that you don't need to modify intervals in-place. You can make a new array and return it.

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0)
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
