package majorityElement

//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than ⌊n / 2⌋ times.
//You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {
	countNum := make(map[int]int)
	halfLen := len(nums) / 2

	for _, num := range nums {
		countNum[num]++
	}

	for k, v := range countNum {
		if v > halfLen {
			return k
		}
	}

	return 0
}
