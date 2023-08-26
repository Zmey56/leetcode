package main

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
		{"abcde", "ebcda"},
		{"aeiou", "uoiea"},
	}

	for _, tc := range testCases {
		result := reverseVowels(tc.input)
		if result != tc.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", tc.input, tc.expected, result)
		}
	}
}
