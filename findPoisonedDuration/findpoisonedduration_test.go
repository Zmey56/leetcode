package findPoisonedDuration

import "testing"

// Test for function findPoisonedDuration

func Test_findPoisonedDuration(t *testing.T) {
	type args struct {
		timeSeries []int
		duration   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test case 1",
			args{[]int{1, 4}, 2},
			4,
		},
		{
			"Test case 2",
			args{[]int{1, 2}, 2},
			3,
		},
		{
			"Test case 3",
			args{[]int{1, 2, 3, 4, 5}, 5},
			9,
		},
		{
			"Test case 4",
			args{[]int{1, 2, 3, 4, 5}, 1},
			5,
		},
		{
			"Test case 5",
			args{[]int{1, 2, 3, 4, 5}, 3},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
