package main

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}

	for _, test := range tests {
		got := longestSubarray(test.nums)
		if got != test.want {
			t.Errorf("longestSubarray(%v) = %d; want %d", test.nums, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{3, 5, 5},
		{10, 7, 10},
		{0, -1, 0},
	}

	for _, test := range tests {
		got := max(test.a, test.b)
		if got != test.want {
			t.Errorf("max(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
		}
	}
}
