package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		output [][]int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix: [][]int{{5}},
			output: [][]int{{5}},
		},
		{
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			output: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, test := range testCases {
		rotate(test.matrix)
		fmt.Println(test.matrix)
		if !equalArray(test.matrix, test.output) {
			t.Errorf("rotate(%v) = %v; want %v", test.matrix, test.matrix, test.output)
		}
	}
}

func equalArray(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
