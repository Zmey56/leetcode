package maxCount

//You are given an m x n matrix M initialized with all 0's and an array of operations ops,
//where ops[i] = [ai, bi] means M[x][y] should be incremented by one for all 0 <= x < ai and 0 <= y < bi.
//
//Count and return the number of maximum integers in the matrix after performing all the operations.

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minA, minB := ops[0][0], ops[0][1]
	for _, op := range ops {
		if op[0] < minA {
			minA = op[0]
		}
		if op[1] < minB {
			minB = op[1]
		}
	}
	return minA * minB
}
