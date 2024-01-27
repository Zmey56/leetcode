package combinationSum2

//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in
//candidates where the candidate numbers sum to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note: The solution set must not contain duplicate combinations.

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	backtrack(&result, path, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, path []int, candidates []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrack(result, path, candidates, target-candidates[i], i+1)
		path = path[:len(path)-1]
	}
}
