package findLengthOfLCIS

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{
				nums: []int{1, 3, 5, 7},
			},
			want: 4,
		},
		{
			name: "Test case 4",
			args: args{
				nums: []int{1, 3, 5, 4, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "Test case 5",
			args: args{
				nums: []int{1, 3, 5, 4, 2, 3, 4, 5, 6},
			},
			want: 5,
		},
		{
			name: "Test case 6",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Test case 7",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "Test case 8",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "Test case 9",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
