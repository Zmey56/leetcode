package isValidSudokuV2

import "sync"

// 36. Valid Sudoku

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
	validMap := make(map[byte]struct{})
	//Check rows
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, exists := validMap[board[i][j]]; exists {
					return false
				}
				validMap[board[i][j]] = struct{}{}
			}
		}
		validMap = make(map[byte]struct{})
	}
	//Check columns
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if _, exists := validMap[board[i][j]]; exists {
					return false
				}
				validMap[board[i][j]] = struct{}{}
			}
		}
		validMap = make(map[byte]struct{})
	}
	// Check 3x3 sub-boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					row := boxRow*3 + i
					col := boxCol*3 + j
					if board[row][col] != '.' {
						if _, exists := validMap[board[row][col]]; exists {
							return false
						}
						validMap[board[row][col]] = struct{}{}
					}
				}
			}
			validMap = make(map[byte]struct{})
		}
	}
	return true

}

func isValidSudokuV2(board [][]byte) bool {
	var wg sync.WaitGroup
	errChan := make(chan bool, 27)

	// Check all rows concurrently
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			validMap := make(map[byte]struct{})
			for j := 0; j < 9; j++ {
				if board[row][j] != '.' {
					if _, exists := validMap[board[row][j]]; exists {
						errChan <- false
						return
					}
					validMap[board[row][j]] = struct{}{}
				}
			}
		}(i)
	}

	// Check all columns concurrently
	for j := 0; j < 9; j++ {
		wg.Add(1)
		go func(col int) {
			defer wg.Done()
			validMap := make(map[byte]struct{})
			for i := 0; i < 9; i++ {
				if board[i][col] != '.' {
					if _, exists := validMap[board[i][col]]; exists {
						errChan <- false
						return
					}
					validMap[board[i][col]] = struct{}{}
				}
			}
		}(j)
	}

	// Check all 3x3 sub-boxes concurrently
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			wg.Add(1)
			go func(br, bc int) {
				defer wg.Done()
				validMap := make(map[byte]struct{})
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						row := br*3 + i
						col := bc*3 + j
						if board[row][col] != '.' {
							if _, exists := validMap[board[row][col]]; exists {
								errChan <- false
								return
							}
							validMap[board[row][col]] = struct{}{}
						}
					}
				}
			}(boxRow, boxCol)
		}
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for valid := range errChan {
		if !valid {
			return false
		}
	}

	return true

}
