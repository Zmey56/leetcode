package thirdMax

import (
	"log"
	"sort"
)

// Given an integer array nums, return the third distinct maximum number in this array.
//If the third maximum does not exist, return the maximum number.

func thirdMax(nums []int) int {
	m := make(map[int]int)
	var res []int
	for _, v := range nums {
		m[v]++
	}
	for k, _ := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	log.Println("RES", res)
	if len(res) < 3 {
		return res[len(res)-1]
	}
	return res[len(res)-3]

}
