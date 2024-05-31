package findShortestSubArray

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 2, 3, 1}}, 2},
		{"2", args{[]int{1, 2, 2, 3, 1, 4, 2}}, 6},
		{"3", args{[]int{1, 2, 2, 3, 1, 4, 2, 2}}, 7},
		{"4", args{[]int{1, 2, 2, 3, 1, 4, 2, 2, 2}}, 7},
		{"5", args{[]int{1, 2, 2, 3, 1, 4, 2, 2, 2, 2}}, 8},
		{"6", args{[]int{1, 2, 2, 3, 1, 4, 2, 2, 2, 2, 2}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
