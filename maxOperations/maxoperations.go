package main

func maxOperations(nums []int, k int) int {
	count := 0
	visited := make(map[int]int)

	for _, num := range nums {
		if visited[k-num] > 0 {
			count++
			visited[k-num]--
		} else {
			visited[num]++
		}
	}

	return count
}
