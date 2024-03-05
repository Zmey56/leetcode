package containsNearbyDuplicate

import "log"

//Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such
//that nums[i] == nums[j] and abs(i - j) <= k.

func containsNearbyDuplicate(nums []int, k int) bool {
	allValues := make(map[int]int)
	for i, v := range nums {
		log.Println(i, v, " - ", allValues[v])
		if _, ok := allValues[v]; ok {
			log.Println("ok", i, v)
			if i-allValues[v] <= k {
				return true
			}
		}
		allValues[v] = i
	}
	return false
}
