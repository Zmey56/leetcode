package main

import "testing"

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := findMaxAverage(tc.nums, tc.k)
			if result != tc.expected {
				t.Errorf("For nums=%v and k=%d, expected %.5f, but got %.5f", tc.nums, tc.k, tc.expected, result)
			}
		})
	}
}
