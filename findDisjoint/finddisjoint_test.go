package main

import (
	"reflect"
	"testing"
)

func TestFindDisjoint(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, test := range tests {
		result := findDisjoint(test.nums1, test.nums2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums1 %v and nums2 %v, expected %v, but got %v", test.nums1, test.nums2, test.expected, result)
		}
	}
}
