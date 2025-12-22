package gameOfLife

import "testing"

func TestGameOfLife(t *testing.T) {
	testCases := []struct {
		board    [][]int
		expected [][]int
	}{
		{
			board: [][]int{
				[]int{0, 1, 0, 0},
				[]int{1, 1, 1, 0},
				[]int{0, 1, 0, 0},
				[]int{1, 1, 0, 0},
			},
			expected: [][]int{
				[]int{0, 0, 1, 0},
				[]int{0, 0, 1, 0},
				[]int{1, 0, 1, 0},
				[]int{0, 1, 1, 0},
			},
		},
		{
			board: [][]int{
				[]int{1, 1},
				[]int{1, 0},
			},
			expected: [][]int{
				[]int{1, 1},
				[]int{1, 1},
			},
		},
	}

	for testNumber, testCase := range testCases {
		gameOfLife(testCase.board)
		if !equalBoards(testCase.board, testCase.expected) {
			t.Errorf("For test #%d, expected %v, got %v", testNumber, testCase.expected, testCase.board)
		}
	}
}

func equalBoards(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
