package findSubstringV2

//You are given a string s and an array of strings words. All the strings of words are of the same length.
//
//A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.
//
//For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
//Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

func findSubstring(s string, words []string) []int {
	wordFrequency := make(map[string]int)

	for _, word := range words {
		wordFrequency[word]++
	}

	var res []int

	length := len(words[0])

	for i := 0; i < len(s)-length*len(words)+1; i++ {
		seen := make(map[string]int)

		for j := 0; j < len(words); j++ {
			nextIndex := i + j*length
			word := s[nextIndex : nextIndex+length]

			if _, ok := wordFrequency[word]; !ok {
				break
			}

			seen[word]++

			seenFrequency, _ := seen[word]
			originFrequency, _ := wordFrequency[word]

			if seenFrequency > originFrequency {
				break
			}

			if j+1 == len(words) {
				res = append(res, i)
			}
		}
	}

	return res
}
