package containsDuplicate

//Given an integer array nums, return true if any value appears at least twice in the array,
//and return false if every element is distinct.

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
