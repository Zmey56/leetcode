package spiralOrder

import "testing"

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 4, 3},
		},
	}

	for _, tc := range testCases {
		got := spiralOrder(tc.matrix)
		if !equal(got, tc.want) {
			t.Errorf("spiralOrder(%v) = %v; want %v", tc.matrix, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
