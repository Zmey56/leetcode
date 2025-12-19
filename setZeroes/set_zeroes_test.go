package setZeroes

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		output [][]int
	}{
		{
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			output: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			output: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, testCase := range testCases {
		setZeroes(testCase.matrix)
		if !equalMatrix(testCase.matrix, testCase.output) {
			t.Errorf("setZeroes(%v) = %v; want %v", testCase.matrix, testCase.matrix, testCase.output)
		}

	}
}

func equalMatrix(a, b [][]int) bool {
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
