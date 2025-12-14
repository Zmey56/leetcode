package findSubstringV2

//You are given a string s and an array of strings words. All the strings of words are of the same length.
//
//A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.
//
//For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
//Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	if len(s) < totalLen {
		return []int{}
	}

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	var res []int

	// Process each offset (0 to wordLen-1)
	for offset := 0; offset < wordLen; offset++ {
		left := offset
		count := 0
		seen := make(map[string]int)

		for right := offset; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]

			if freq, exists := wordFreq[word]; exists {
				seen[word]++
				count++

				// Shrink window if we have too many of this word
				for seen[word] > freq {
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					count--
					left += wordLen
				}

				// Check if we have a valid window
				if count == wordCount {
					res = append(res, left)
					// Slide window by one word
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					count--
					left += wordLen
				}
			} else {
				// Reset window if word not in dictionary
				seen = make(map[string]int)
				count = 0
				left = right + wordLen
			}
		}
	}

	return res
}
