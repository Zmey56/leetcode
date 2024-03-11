package missingNumber

import "testing"

// Test for missingNumber
func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}
	for _, test := range tests {
		if got := missingNumber(test.nums); got != test.want {
			t.Errorf("missingNumber(%v) = %v; want %v", test.nums, got, test.want)
		}
	}
}
