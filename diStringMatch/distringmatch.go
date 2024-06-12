package diStringMatch

// A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:
//
//s[i] == 'I' if perm[i] < perm[i + 1], and
//s[i] == 'D' if perm[i] > perm[i + 1].
//Given a string s, reconstruct the permutation perm and return it. If there are multiple valid permutations perm, return any of them.

func diStringMatch(s string) []int {
	high := len(s)
	low := 0
	result := make([]int, len(s)+1)
	for i, v := range s {
		if v == 'I' {
			result[i] = low
			low++
		} else {
			result[i] = high
			high--
		}
	}
	result[len(s)] = low
	return result
}
