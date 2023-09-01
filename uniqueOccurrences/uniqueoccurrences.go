package main

func uniqueOccurrences(arr []int) bool {
	result := make(map[int]int)

	for _, num := range arr {
		result[num]++
	}

	counts := make(map[int]bool)

	for _, count := range result {
		if counts[count] {
			return false
		}
		counts[count] = true
	}

	return true
}
