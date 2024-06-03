package surfaceArea

import "testing"

func Test_surfaceArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				grid: [][]int{
					{2},
				},
			},
			want: 10,
		},
		{
			name: "Test case 2",
			args: args{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: 34,
		},
		{
			name: "Test case 3",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 2},
				},
			},
			want: 16,
		},
		{
			name: "Test case 4",
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: 32,
		},
		{
			name: "Test case 5",
			args: args{
				grid: [][]int{
					{2, 2, 2},
					{2, 1, 2},
					{2, 2, 2},
				},
			},
			want: 46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := surfaceArea(tt.args.grid); got != tt.want {
				t.Errorf("surfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
