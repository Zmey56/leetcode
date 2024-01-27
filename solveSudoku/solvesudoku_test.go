package solveSudoku

import "testing"

// Test for function solveSudoku

func TestSolveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte // want is the expected result from calling solveSudoku
	}{
		{
			name: "Test #1",
			args: args{
				board: [][]byte{
					[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: [][]byte{
				[]byte{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				[]byte{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				[]byte{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				[]byte{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				[]byte{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				[]byte{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				[]byte{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				[]byte{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				[]byte{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			if !compare(tt.args.board, tt.want) {
				t.Errorf("solveSudoku() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}

// compare compares two 2D slices of bytes
func compare(a, b [][]byte) bool {
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
