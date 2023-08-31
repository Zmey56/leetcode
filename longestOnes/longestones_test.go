package main

import (
	"testing"
)

func TestMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, test := range tests {
		got := longestOnes(test.nums, test.k)
		if got != test.want {
			t.Errorf("maxConsecutiveOnes(%v, %d) = %d; want %d", test.nums, test.k, got, test.want)
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
