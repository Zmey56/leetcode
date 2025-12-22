package gameOfLife

//289. Game of Life

//According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the
//British mathematician John Horton Conway in 1970."
//
//The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead
//(represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following
//four rules (taken from the above Wikipedia article):
//
//Any live cell with fewer than two live neighbors dies as if caused by under-population.
//Any live cell with two or three live neighbors lives on to the next generation.
//Any live cell with more than three live neighbors dies, as if by over-population.
//Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
//The next state of the board is determined by applying the above rules simultaneously to every cell in the current state
//of the m x n grid board. In this process, births and deaths occur simultaneously.
//
//Given the current state of the board, update the board to reflect its next state.
//
//Note that you do not need to return anything.

//Follow up:

//Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells
//first and then use their updated values to update other cells.
//In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause
//problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would
//you address these problems?

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	// Directions: 8 neighbors
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					// 1 or 2 means originally alive
					if board[ni][nj] == 1 || board[ni][nj] == 2 {
						liveNeighbors++
					}
				}
			}

			if board[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[i][j] = 2 // alive -> dead
			}
			if board[i][j] == 0 && liveNeighbors == 3 {
				board[i][j] = 3 // dead -> alive
			}
		}
	}

	// Convert temporary states to final states
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}
