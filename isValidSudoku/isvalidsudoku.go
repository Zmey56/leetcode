package isValidSudoku

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
//Each row must contain the digits 1-9 without repetition.
//Each column must contain the digits 1-9 without repetition.
//Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//Note:
//
//A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//Only the filled cells need to be validated according to the mentioned rules.

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]int
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				k := i/3*3 + j/3
				if row[i][num] == 1 || col[j][num] == 1 || box[k][num] == 1 {
					return false
				}
				row[i][num] = 1
				col[j][num] = 1
				box[k][num] = 1
			}
		}
	}
	return true
}
