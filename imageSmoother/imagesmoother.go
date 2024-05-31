package imageSmoother

//An image smoother is a filter of the size 3 x 3 that can be applied to each cell of an image by rounding down the average of
//the cell and the eight surrounding cells (i.e., the average of the nine cells in the blue smoother).
//If one or more of the surrounding cells of a cell is not present, we do not consider it in the average (i.e., the average of the four cells in the red smoother).

func imageSmoother(img [][]int) [][]int {
	rows, cols := len(img), len(img[0])
	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[i][j] = smooth(img, i, j)
		}
	}
	return res

}

// smooth calculates the average of the cell and the eight surrounding cells
func smooth(img [][]int, i, j int) int {
	rows, cols := len(img), len(img[0])
	sum, count := 0, 0
	for r := i - 1; r <= i+1; r++ {
		for c := j - 1; c <= j+1; c++ {
			if r >= 0 && r < rows && c >= 0 && c < cols {
				sum += img[r][c]
				count++
			}
		}
	}
	return sum / count
}
