package surfaceArea

//You are given an n x n grid where you have placed some 1 x 1 x 1 cubes. Each value v = grid[i][j] represents a tower of
// v cubes placed on top of cell (i, j).
//
//After placing these cubes, you have decided to glue any directly adjacent cubes to each other, forming several irregular 3D shapes.
//
//Return the total surface area of the resulting shapes.
//
//Note: The bottom face of each shape counts toward its surface area.

func surfaceArea(grid [][]int) int {
	n := len(grid)
	var area int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				area += 2
				area += 4 * grid[i][j]
				if i > 0 {
					area -= 2 * min(grid[i][j], grid[i-1][j])
				}
				if j > 0 {
					area -= 2 * min(grid[i][j], grid[i][j-1])
				}
			}
		}
	}
	return area

}
