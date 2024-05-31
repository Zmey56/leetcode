package searchTwo

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5}, 4},
		{"2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1}, 0},
		{"3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9}, 8},
		{"4", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 10}, -1},
		{"5", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}

}
