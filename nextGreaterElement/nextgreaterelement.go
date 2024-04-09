package nextGreaterElement

//The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.
//
//You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
//
//For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2.
//If there is no next greater element, then the answer for this query is -1.
//
//Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	for _, num1 := range nums1 {
		var nextGreter = -1
		for i, num2 := range nums2 {
			if num1 == num2 {
				for j := i + 1; j < len(nums2); j++ {
					if nums2[j] > num1 {
						nextGreter = nums2[j]
						break
					}
				}
			}
		}
		result = append(result, nextGreter)
	}

	return result
}
