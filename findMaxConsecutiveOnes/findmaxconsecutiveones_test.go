package findMaxConsecutiveOnes

import "testing"

// Test for function findMaxConsecutiveOnes

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test case 1",
			args{[]int{1, 1, 0, 1, 1, 1}},
			3,
		},
		{
			"Test case 2",
			args{[]int{1, 0, 1, 1, 0, 1}},
			2,
		},
		{
			"Test case 3",
			args{[]int{1, 1, 1, 1, 1, 1}},
			6,
		},
		{
			"Test case 4",
			args{[]int{0, 0, 0, 0, 0, 0}},
			0,
		},
		{
			"Test case 5",
			args{[]int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}},
			6,
		},
		{
			"Test case 6",
			args{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
