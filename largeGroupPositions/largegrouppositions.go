package largeGroupPositions

//In a string s of lowercase letters, these letters form consecutive groups of the same character.
//
//For example, a string like s = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z", and "yy".
//
//A group is identified by an interval [start, end], where start and end denote the start and end indices (inclusive) of the group.
//In the above example, "xxxx" has the interval [3,6].
//
//A group is considered large if it has 3 or more characters.
//
//Return the intervals of every large group sorted in increasing order by start index.

type LargeGroupPositions struct {
	begin  int
	end    int
	lenInt int
}

func largeGroupPositions(s string) [][]int {
	result := []LargeGroupPositions{}
	prev := s[0]
	count := 1

	for i := 1; i <= len(s)-1; i++ {
		if s[i] != prev {
			if count >= 3 {
				result = append(result, LargeGroupPositions{begin: i - count, end: i - 1, lenInt: count})
				prev = s[i]
				count = 0
			} else {
				prev = s[i]
				count = 0
			}
		}
		count++
	}

	if count >= 3 {
		result = append(result, LargeGroupPositions{begin: len(s) - count, end: len(s) - 1, lenInt: count})
	}

	//sort.Slice(result, func(i, j int) bool {
	//	return result[i].lenInt < result[j].lenInt
	//})

	res := [][]int{}
	//sort LargeGroupPositions by lenInt
	for _, v := range result {
		if v.lenInt >= 3 {
			res = append(res, []int{v.begin, v.end})
		}
	}

	return res
}

func largeGroupPositionsVer2(s string) [][]int {
	result := [][]int{}
	prev := s[0]
	count := 1
	start := 0

	for i := 1; i < len(s); i++ {
		if s[i] != prev {
			if count >= 3 {
				result = append(result, []int{start, i - 1})
			}
			start = i
			count = 1
			prev = s[i]
		} else {
			count++
		}
	}

	if count >= 3 {
		result = append(result, []int{start, len(s) - 1})
	}

	return result
}
