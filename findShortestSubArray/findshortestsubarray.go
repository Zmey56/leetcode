package findShortestSubArray

//Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.
//
//Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

func findShortestSubArray(nums []int) int {
	degree := 0
	count := make(map[int]int)
	first := make(map[int]int)
	res := 0
	for i, num := range nums {
		if _, ok := first[num]; !ok {
			first[num] = i
		}
		count[num]++
		if count[num] > degree {
			degree = count[num]
			res = i - first[num] + 1
		} else if count[num] == degree {
			res = min(res, i-first[num]+1)
		}
	}
	return res

}
