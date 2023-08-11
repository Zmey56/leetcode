package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input:    "",
			expected: []string{},
		}, {
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actual := letterCombinations(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected: %s, but got: %s", testCase.expected, actual)
			}
		})
	}
}
