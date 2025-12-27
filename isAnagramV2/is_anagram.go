package isAnagramV2

// 242. Valid Anagram

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	//mapT := make(map[rune]int)
	for _, letter := range s {
		mapS[letter]++
	}

	for _, letter := range t {
		if _, ok := mapS[letter]; !ok {
			return false
		}
		mapS[letter]--
		if mapS[letter] == 0 {
			delete(mapS, letter)
		}
	}

	if len(mapS) > 0 {
		return false
	}

	return true

}
