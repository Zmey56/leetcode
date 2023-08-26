package main

import "testing"

func TestCompress(t *testing.T) {
	testCases := []struct {
		chars    []byte
		expected int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	}

	for _, tc := range testCases {
		result := compress(tc.chars)
		if result != tc.expected {
			t.Errorf("For %v, expected length %v but got %v", tc.chars, tc.expected, result)
		}
	}
}
