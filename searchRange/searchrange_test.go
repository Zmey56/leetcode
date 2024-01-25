package searchRange

import (
	"reflect"
	"testing"
)

//Test for function searchRange

func TestSearchRange(t *testing.T) {
	test := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   []int{4, 4},
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 13,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -2,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -1,
			want:   []int{0, 0},
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 0,
			want:   []int{1, 1},
		},
	}
	for _, test := range test {
		got := searchRange(test.nums, test.target)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("searchRange(%v, %d) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
