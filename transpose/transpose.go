package transpose

//Given a 2D integer array matrix, return the transpose of matrix.
//
//The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.

func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	transposed := make([][]int, m)
	for i := 0; i < m; i++ {
		transposed[i] = make([]int, n)
		for j := 0; j < n; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}

	return transposed

}
