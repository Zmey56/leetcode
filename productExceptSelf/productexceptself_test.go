package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		slice    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range testCases {
		result := productExceptSelf(test.slice)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.slice, test.expected, result)
		}
	}
}
