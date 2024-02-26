package getRow

//Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]

		}

	}
	return res[rowIndex]
}
