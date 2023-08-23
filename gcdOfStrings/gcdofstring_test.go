package main

import (
	"testing"
)

func TestGCDOfStrings(t *testing.T) {
	testCases := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"AAAAA", "A", "A"},
	}

	for _, tc := range testCases {
		t.Run(tc.str1+"_"+tc.str2, func(t *testing.T) {
			result := gcdOfStrings(tc.str1, tc.str2)
			if result != tc.expected {
				t.Errorf("For '%s' and '%s', expected '%s' but got '%s'", tc.str1, tc.str2, tc.expected, result)
			}
		})
	}
}
