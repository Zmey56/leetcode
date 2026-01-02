package twoSum

import "sort"

//1. Two Sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	mapWithValue := make(map[int]int)
	resultIndex := make([]int, 2)
	for i, v := range nums {
		if _, ok := mapWithValue[v]; ok {
			if v+v == target {
				return []int{mapWithValue[v], i}
			}
		}
		mapWithValue[v] = i
	}

	for key, value := range mapWithValue {
		secondValue := target - key
		if secondValueIndex, ok := mapWithValue[secondValue]; ok && secondValueIndex != value {
			resultIndex[0] = value
			resultIndex[1] = secondValueIndex
			sort.Ints(resultIndex)
			return resultIndex
		}
	}
	return resultIndex
}
