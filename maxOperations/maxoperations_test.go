package main

import "testing"

func TestMaxOperations(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{1, 1, 1, 1}, 2, 2},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := maxOperations(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("For nums=%v and k=%d, expected %d, but got %d", tc.nums, tc.k, tc.expected, result)
			}
		})
	}
}
