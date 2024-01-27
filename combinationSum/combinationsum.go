package combinationSum

//Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of
//candidates where the chosen numbers sum to target. You may return the combinations in any order.
//
//The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
//frequency
// of at least one of the chosen numbers is different.
//
//The test cases are generated such that the number of unique combinations that sum up to target is less than
//150 combinations for the given input.

func combinationSum(candidate []int, target int) [][]int {
	var result [][]int
	var path []int
	backtrack(&result, path, candidate, target, 0)
	return result
}

func backtrack(result *[][]int, path []int, candidate []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(candidate); i++ {
		path = append(path, candidate[i])
		backtrack(result, path, candidate, target-candidate[i], i)
		path = path[:len(path)-1]
	}
}
