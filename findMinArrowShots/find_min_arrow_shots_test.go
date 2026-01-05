package findMinArrowShots

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "Example 1",
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
		{
			name:   "Example 2",
			points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:   4,
		},
		{
			name:   "Example 3",
			points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:   2,
		},
	}

	for _, tt := range tests {
		if got := findMinArrowShots(tt.points); got != tt.want {
			t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
		}
	}
}
