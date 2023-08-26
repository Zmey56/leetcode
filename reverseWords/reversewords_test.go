package main

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for _, test := range tests {
		output := reverseWords(test.input)
		if output != test.expected {
			t.Errorf("For input: %s, expected: %s, but got: %s", test.input, test.expected, output)
		}
	}
}
