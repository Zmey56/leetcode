package main

import "testing"

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}

	for _, test := range tests {
		result := pivotIndex(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d, but got %d", test.nums, test.expected, result)
		}
	}
}
