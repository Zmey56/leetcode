package orangesRotting

// You are given an m x n grid where each cell can have one of three values:
//
//0 representing an empty cell,
//1 representing a fresh orange, or
//2 representing a rotten orange.
//Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
//
//Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	type cell struct{ r, c int }

	q := make([]cell, 0)
	fresh := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch grid[r][c] {
			case 2:
				q = append(q, cell{r, c})
			case 1:
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0

	for len(q) > 0 && fresh > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[i]
			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					q = append(q, cell{nr, nc})
				}
			}
		}
		q = q[size:]
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}
