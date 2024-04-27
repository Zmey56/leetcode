package matrixReshape

// In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
//
//You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.
//
//The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
//
//If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = mat[i/n][i%n]
	}

	return res

}
