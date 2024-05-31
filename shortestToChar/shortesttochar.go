package shortestToChar

//Given a string s and a character c that occurs in s, return an array of integers answer where answer.length == s.length
//and answer[i] is the distance from index i to the closest occurrence of character c in s.
//
//The distance between two indices i and j is abs(i - j), where abs is the absolute value function.

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	pos := -n

	for i := 0; i < n; i++ {
		if s[i] == c {
			pos = i
		}
		res[i] = i - pos
	}

	for i := pos - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		if pos-i < res[i] {
			res[i] = pos - i
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
