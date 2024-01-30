package combinationSum2

import (
	"sort"
)

//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in
//candidates where the candidate numbers sum to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note: The solution set must not contain duplicate combinations.

func combinationUtil(cc *[]int, candidates []int, idx, currentSum, target, cclen int, res *[][]int) {
	if currentSum == target {
		combi := make([]int, cclen)
		copy(combi, *cc)

		*res = append(*res, combi)

		return
	}

	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}

		if currentSum+candidates[i] > target {
			break
		}
		(*cc)[cclen] = candidates[i]
		combinationUtil(cc, candidates, i+1, currentSum+candidates[i], target, cclen+1, res)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	cc := make([]int, target)
	combinationUtil(&cc, candidates, 0, 0, target, 0, &res)

	return res
}
