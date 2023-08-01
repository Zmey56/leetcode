package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))

	nums = []int{0, 0, 0}
	target = 1
	fmt.Println(threeSumClosest(nums, target))

	nums = []int{0, 1, 2}
	target = 3
	fmt.Println(threeSumClosest(nums, target))

	nums = []int{1, 1, 1, 0}
	target = -100
	fmt.Println(threeSumClosest(nums, target))

}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	closestSum := math.MaxInt32

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}

			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				// If we find an exact match, we can directly return the target
				return target
			}
		}
	}

	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
