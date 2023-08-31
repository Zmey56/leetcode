package main

import "testing"

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain []int
		want int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, test := range tests {
		got := largestAltitude(test.gain)
		if got != test.want {
			t.Errorf("largestAltitude(%v) = %d; want %d", test.gain, got, test.want)
		}
	}
}
