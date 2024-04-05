package islandPerimeter

import "testing"

// Test for function islandPerimeter

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test case 1",
			args{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}},
			16,
		},
		{
			"Test case 2",
			args{[][]int{{1}}},
			4,
		},
		{
			"Test case 3",
			args{[][]int{{1, 0}}},
			4,
		},
		{
			"Test case 4",
			args{[][]int{{1, 1}}},
			6,
		},
		{
			"Test case 5",
			args{[][]int{{0, 1}, {1, 1}}},
			8,
		},
		{
			"Test case 6",
			args{[][]int{{1, 1}, {1, 1}}},
			8,
		},
		{
			"Test case 7",
			args{[][]int{{0, 0}, {0, 0}}},
			0,
		},
		{
			"Test case 8",
			args{[][]int{{1, 0}, {0, 0}}},
			4,
		},
		{
			"Test case 9",
			args{[][]int{{1, 1}, {0, 0}}},
			6,
		},
		{
			"Test case 10",
			args{[][]int{{1, 1}, {1, 0}}},
			8,
		},
		{
			"Test case 11",
			args{[][]int{{1, 1}, {1, 1}}},
			8,
		},
		{
			"Test case 12",
			args{[][]int{{0, 1}, {1, 0}}},
			8,
		},
		{
			"Test case 13",
			args{[][]int{{0, 1}, {0, 1}}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
