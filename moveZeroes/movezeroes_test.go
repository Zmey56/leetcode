package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		// Test cases with mixed numbers and zeros
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 0, 1, 2}, []int{1, 2, 0, 0, 0}},
		{[]int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},

		// Test cases with no zeros
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{-1, -2, -3, -4, -5}, []int{-1, -2, -3, -4, -5}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},

		// Edge cases
		{[]int{}, []int{}},   // Empty input
		{[]int{0}, []int{0}}, // Single zero
		{[]int{1}, []int{1}}, // Single non-zero
	}

	for _, test := range tests {
		// Make a copy of the input to preserve the original
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		moveZeroes(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", inputCopy, test.expected, test.input)
		}
	}
}
