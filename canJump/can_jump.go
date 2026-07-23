package main

// 55. Jump Game

// You are given an integer array nums. You are initially positioned at the array's
//  first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.
func canJump(nums []int) bool {
	furthestReach := 0

	for i := 0; i <= len(nums); i++ {
		if i > furthestReach {
			return false
		}

		currentReach := i + nums[i]

		if currentReach > furthestReach {
			furthestReach = currentReach
		}

		if furthestReach >= len(nums)-1 {
			return true
		}
	}

	return true
}
