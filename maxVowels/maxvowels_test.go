package main

import "testing"

func TestMaxVowels(t *testing.T) {
	testCases := []struct {
		s        string
		k        int
		expected int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := maxVowels(tc.s, tc.k)
			if result != tc.expected {
				t.Errorf("For s=%s and k=%d, expected %d, but got %d", tc.s, tc.k, tc.expected, result)
			}
		})
	}
}
