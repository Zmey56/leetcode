package longestConsecutive

// 128. Longest Consecutive Sequence

//Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
//You must write an algorithm that runs in O(n) time.

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		set[v] = true
	}

	maxLen := 0

	for n := range set {
		if !set[n-1] {
			current := n
			length := 1

			for set[current+1] {
				current++
				length++
			}

			if length > maxLen {
				maxLen = length
			}

		}
	}

	return maxLen
}
