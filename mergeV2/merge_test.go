package mergeV2

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{4, 7}, {1, 4}},
			expected: [][]int{{1, 7}},
		},
	}

	for _, tt := range tests {
		res := merge(tt.input)
		if reflect.DeepEqual(res, tt.expected) == false {
			t.Errorf("merge(%v) = %v, want %v", tt.input, res, tt.expected)
		}
	}
}
