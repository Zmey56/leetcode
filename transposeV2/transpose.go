package main

// 867. Transpose Matrix

// Given a 2D integer array matrix, return the transpose of matrix.

// The transpose of a matrix is the matrix flipped over its main diagonal,
//
//	switching the matrix's row and column indices.
func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, m)
	}
	for i, row := range matrix {
		for j, val := range row {
			out[j][i] = val
		}
	}
	return out
}
