package main

import "fmt"

// 88. Merge Sorted Array

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
//  and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
//
//	To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements
//	that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Pointers for nums1, nums2, and the very end of the merged array
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// While there are elements to compare in both arrays
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	fmt.Println(nums1)

	// If there are leftover elements in nums2, copy them over.
	// (We don't need to check leftover nums1 elements because they are already in place).
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
