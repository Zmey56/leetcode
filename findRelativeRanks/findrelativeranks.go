package findRelativeRanks

import (
	"sort"
	"strconv"
)

//You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.
//
//The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on.
//The placement of each athlete determines their rank:
//
//The 1st place athlete's rank is "Gold Medal".
//The 2nd place athlete's rank is "Silver Medal".
//The 3rd place athlete's rank is "Bronze Medal".
//For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").
//Return an array answer of size n where answer[i] is the rank of the ith athlete.

func findRelativeRanks(score []int) []string {
	var temp []int
	mp := map[int]int{}
	var ans []string

	temp = append(temp, score...)

	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	for i := range temp {
		mp[temp[i]] = i + 1
	}

	for i := range score {
		if mp[score[i]] == 1 {
			ans = append(ans, "Gold Medal")
		} else if mp[score[i]] == 2 {
			ans = append(ans, "Silver Medal")
		} else if mp[score[i]] == 3 {
			ans = append(ans, "Bronze Medal")
		} else {
			num := strconv.Itoa(mp[score[i]])
			ans = append(ans, num)
		}
	}

	return ans

}
