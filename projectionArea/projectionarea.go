package projectionArea

//You are given an n x n grid where we place some 1 x 1 x 1 cubes that are axis-aligned with the x, y, and z axes.
//
//Each value v = grid[i][j] represents a tower of v cubes placed on top of the cell (i, j).
//
//We view the projection of these cubes onto the xy, yz, and zx planes.
//
//A projection is like a shadow, that maps our 3-dimensional figure to a 2-dimensional plane.
//We are viewing the "shadow" when looking at the cubes from the top, the front, and the side.
//
//Return the total area of all three projections.

func projectionArea(grid [][]int) int {
	var xy, yz, zx int
	n := len(grid)
	for i := 0; i < n; i++ {
		var maxRow, maxCol int
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				xy++
			}
			maxRow = max(maxRow, grid[i][j])
			maxCol = max(maxCol, grid[j][i])
		}
		yz += maxRow
		zx += maxCol
	}
	return xy + yz + zx

}
