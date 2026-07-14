package main

//35. Search Insert Position

// Given a sorted array of distinct integers and a target value,
// return the index if the target is found. If not, return the index where it would be
// if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

func searchInsert(nums []int, target int) int {
	index := 0
	for _, j := range nums {
		if j >= target {
			return index
		}
		index++
	}
	return index
}
