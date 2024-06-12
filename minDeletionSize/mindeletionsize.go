package minDeletionSize

import "sync"

//You are given an array of n strings strs, all of the same length.
//
//The strings can be arranged such that there is one on each line, making a grid.
//
//For example, strs = ["abc", "bce", "cae"] can be arranged as follows:
//abc
//bce
//cae
//You want to delete the columns that are not sorted lexicographically. In the above example (0-indexed),
//columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted, while column 1 ('b', 'c', 'a') is not, so you would delete column 1.

func minDeletionSize(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				res++
				break
			}
		}
	}
	return res

}

func minDeletionSize2(strs []string) int {
	var wg sync.WaitGroup
	resChan := make(chan bool, len(strs[0]))

	for i := 0; i < len(strs[0]); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 1; j < len(strs); j++ {
				if strs[j][idx] < strs[j-1][idx] {
					resChan <- true
					return
				}
			}
			resChan <- false
		}(i)
	}

	wg.Wait()
	close(resChan)

	res := 0
	for result := range resChan {
		if result {
			res++
		}
	}

	return res

}

func minDeletionSize3(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		isSorted := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				isSorted = false
				break
			}
		}
		if !isSorted {
			res++
		}
	}
	return res
}
