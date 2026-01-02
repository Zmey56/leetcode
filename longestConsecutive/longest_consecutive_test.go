package longestConsecutive

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
	}

	for _, tt := range tests {
		if got := longestConsecutive(tt.nums); got != tt.want {
			t.Errorf("longestConsecutive(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
