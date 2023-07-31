package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

	height = []int{1, 1}
	fmt.Println(maxArea(height))

}

func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	maxWater := 0

	for left < right {
		width := right - left
		currentWater := width * min(height[left], height[right])
		maxWater = max(maxWater, currentWater)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
