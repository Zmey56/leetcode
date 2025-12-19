package setZeroes

//Test Result
//73. Set Matrix Zeroes

//Hint
//Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
//
//You must do it in place.

func setZeroes(matrix [][]int) {

	colZero := make(map[int]struct{})
	rowZero := make(map[int]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = struct{}{}
				colZero[j] = struct{}{}
			}
		}
	}

	for r := range rowZero {
		for j := 0; j < len(matrix[r]); j++ {
			matrix[r][j] = 0
		}
	}

	for c := range colZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][c] = 0
		}
	}

}
