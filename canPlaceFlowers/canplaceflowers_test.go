package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 0, 0, 0}, 3, true},
		// Add more test cases here
	}

	for i, tc := range testCases {
		result := canPlaceFlowers(tc.flowerbed, tc.n)
		if result != tc.expected {
			t.Errorf("Test case %d failed: expected %v, got %v", i+1, tc.expected, result)
		}
	}
}

func TestMain(m *testing.M) {
	// Run all tests
	m.Run()
}
