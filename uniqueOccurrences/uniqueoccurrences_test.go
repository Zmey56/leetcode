package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		array    []int
		expected bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range testCases {
		result := uniqueOccurrences(test.array)
		if result != test.expected {
			t.Errorf("For nums %v, expected %v, but got %v", test.array, test.expected, result)
		}
	}
}
