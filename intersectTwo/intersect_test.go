package intersectTwo

//Given two integer arrays nums1 and nums2, return an array of their intersection.
//Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

func intersect(nums1 []int, nums2 []int) []int {

	resultMap1 := make(map[int]int)
	resultMap2 := make(map[int]int)

	for _, v := range nums1 {
		resultMap1[v]++
	}

	for _, v := range nums2 {
		resultMap2[v]++
	}

	var result []int

	for k, v := range resultMap1 {
		if resultMap2[k] > 0 {

		}
	}

}
