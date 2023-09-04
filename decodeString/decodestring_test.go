package main

import (
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, test := range tests {
		result := decodeString(test.s)
		if result == test.expected {
			t.Errorf("For asteroid %v, expected %v, but got %v", test.s, test.expected, result)
		}
	}
}
