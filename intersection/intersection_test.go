package intersection

import (
	"reflect"
	"testing"
)

// Test for function intersection
func TestIntersection(t *testing.T) {
	test := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{}},
	}

	for _, test := range test {
		result := intersection(test.nums1, test.nums2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test failed: expected %v, got %v", test.expected, result)
		}
	}
}
