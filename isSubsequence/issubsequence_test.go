package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true}, // An empty string is always a subsequence
		{"abc", "", false},   // A non-empty string can't be a subsequence of an empty string
		{"abc", "abc", true}, // A string is always a subsequence of itself
		{"abc", "ab", false}, // s is longer than t
		{"abc", "ac", false}, // s is not ordered the same way as t
	}

	for _, test := range tests {
		result := isSubsequence(test.s, test.t)
		if result != test.expected {
			t.Errorf("For s=%s, t=%s, expected %v, but got %v", test.s, test.t, test.expected, result)
		}
	}
}
