package projectionArea

import "testing"

func Test_projectionArea(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			want: 17,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{2},
			},
			want: 5,
		},
		{
			name: "Example 3",
			grid: [][]int{
				{1, 0},
				{0, 2},
			},
			want: 8,
		},
		{
			name: "Example 4",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: 14,
		},
		{
			name: "Example 5",
			grid: [][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := projectionArea(tt.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
