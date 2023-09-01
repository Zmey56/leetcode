package main

import "testing"

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"cabbba", "aabbss", false},
	}

	for _, test := range tests {
		result := closeStrings(test.word1, test.word2)
		if result != test.expected {
			t.Errorf("For word1 %v and word2 %v, expected %v, but got %v", test.word1, test.word2, test.expected, result)
		}
	}
}
