package main

func findDisjoint(nums1, nums2 []int) [][]int {
	distinctInNums1 := make(map[int]bool)
	distinctInNums2 := make(map[int]bool)

	for _, num := range nums1 {
		distinctInNums1[num] = true
	}

	for _, num := range nums2 {
		distinctInNums2[num] = true
	}

	answer := [][]int{
		make([]int, 0),
		make([]int, 0),
	}

	for num, _ := range distinctInNums1 {
		if !distinctInNums2[num] {
			answer[0] = append(answer[0], num)
		}
	}

	for num, _ := range distinctInNums2 {
		if !distinctInNums1[num] {
			answer[1] = append(answer[1], num)
		}
	}

	return answer
}
