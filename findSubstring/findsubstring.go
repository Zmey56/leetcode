package findSubstring

//You are given a string s and an array of strings words. All the strings of words are of the same length.
//
//A concatenated substring in s is a substring that contains all the strings of any permutation of words concatenated.
//
//For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are
//all concatenated strings. "acdbef" is not a concatenated substring because it is not the concatenation
//of any permutation of words.
//Return the starting indices of all the concatenated substrings in s. You can return the answer in any order.

func findSubstring(s string, words []string) []int {
	var res []int
	if len(words) == 0 {
		return res
	}
	wordLen := len(words[0])
	wordsLen := len(words)
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}
	for i := 0; i < len(s)-wordLen*wordsLen+1; i++ {
		seen := make(map[string]int)
		j := 0
		for j < wordsLen {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if _, ok := wordsMap[word]; ok {
				seen[word]++
				if seen[word] > wordsMap[word] {
					break
				}
			} else {
				break
			}
			j++
		}
		if j == wordsLen {
			res = append(res, i)
		}
	}
	return res
}
