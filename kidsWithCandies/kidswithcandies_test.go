package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := kidsWithCandies(tc.candies, tc.extraCandies)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For candies %v with extraCandies %d, expected %v but got %v", tc.candies, tc.extraCandies, tc.expected, result)
			}
		})
	}
}
