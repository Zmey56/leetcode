package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, tc := range testCases {
		t.Run(tc.word1+"_"+tc.word2, func(t *testing.T) {
			result := mergeAlternately(tc.word1, tc.word2)
			if result != tc.expected {
				t.Errorf("Expected '%s' but got '%s'", tc.expected, result)
			}
		})
	}
}
