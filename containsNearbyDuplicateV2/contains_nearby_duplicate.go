package containsNearbyDuplicateV2

//219. Contains Duplicate II

//Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such
//that nums[i] == nums[j] and abs(i - j) <= k.

func containsNearbyDuplicate(nums []int, k int) bool {
	lastIndex := make(map[int]int)

	for i, n := range nums {
		if prev, ok := lastIndex[n]; ok && i-prev <= k {
			return true
		}
		lastIndex[n] = i
	}
	return false
}
