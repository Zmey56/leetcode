package main

func equalPairs(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count := 0
			for k := 0; k < len(grid); k++ {
				if grid[i][k] == grid[k][j] {
					count += 1
				} else {
					break
				}
			}
			if count == len(grid) {
				ans += 1
			}
		}

	}
	return ans
}
