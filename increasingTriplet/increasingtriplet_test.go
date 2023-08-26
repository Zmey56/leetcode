package main

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	}

	for _, tc := range testCases {
		result := increasingTriplet(tc.nums)
		if result != tc.expected {
			t.Errorf("For %v, expected %v but got %v", tc.nums, tc.expected, result)
		}
	}
}
