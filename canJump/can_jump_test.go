package main

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want bool
	}{
		{
			name: "Example 1: Can reach the end",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Example 2: Trapped by zero",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Single element array (already at the end)",
			nums: []int{0},
			want: true,
		},
		{
			name: "Stuck at the very beginning",
			nums: []int{0, 2, 3},
			want: false,
		},
		{
			name: "Exactly reaching the last index with exact steps",
			nums: []int{2, 0, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
