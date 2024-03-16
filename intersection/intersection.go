package intersection

//Given two integer arrays nums1 and nums2, return an array of their intersection.
//Each element in the result must be unique and you may return the result in any order.

func intersection(nums1 []int, nums2 []int) []int {
	mapWithValue := make(map[int]bool)

	for _, v := range nums1 {
		mapWithValue[v] = true
	}

	var result []int
	for _, v := range nums2 {
		if mapWithValue[v] {
			result = append(result, v)
			mapWithValue[v] = false
		}
	}

	return result

}
