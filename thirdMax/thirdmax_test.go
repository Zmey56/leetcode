package thirdMax

import "testing"

// Test for function thirdMax
func Test_thirdMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			nums: []int{1, 2},
			want: 2,
		},
		{
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			nums: []int{5, 2, 2},
			want: 5,
		}, {
			nums: []int{1, -2147483648, 2},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		if got := thirdMax(tt.nums); got != tt.want {
			t.Errorf("thirdMax() = %v, want %v", got, tt.want)
		}
	}
}
