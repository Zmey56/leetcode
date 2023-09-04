package main

import "testing"

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}

	for _, test := range tests {
		result := removeStars(test.s)
		if result != test.expected {
			t.Errorf("For nums %v, expected %v, but got %v", test.s, test.expected, result)
		}
	}
}
