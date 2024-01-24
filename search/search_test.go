package search

import "testing"

func TestSearch(t *testing.T) {
	test := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 13,
			want:   -1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -2,
			want:   -1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -1,
			want:   0,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 0,
			want:   1,
		},
	}
	for _, test := range test {
		got := search(test.nums, test.target)
		if got != test.want {
			t.Errorf("search(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}

}
