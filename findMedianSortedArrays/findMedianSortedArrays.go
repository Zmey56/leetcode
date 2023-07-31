package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	avg := len(nums1) % 2

	if avg != 0 {
		idx := int(math.Ceil(float64(len(nums1))/2) - 1)
		return float64(nums1[idx])
	} else {
		fmt.Println(nums1)
		idx1 := int(math.Ceil(float64(len(nums1))/2) - 1)
		idx2 := int(math.Ceil(float64(len(nums1)) / 2))
		return float64(float64(nums1[idx1]+nums1[idx2]) / 2.0)
	}
}
