package wordPatternV2

import "strings"

//290. Word Pattern

//Given a pattern and a string s, find if s follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:
//
//Each letter in pattern maps to exactly one unique word in s.
//Each unique word in s maps to exactly one letter in pattern.
//No two letters map to the same word, and no two words map to the same letter.

func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	dict := make(map[rune]string)
	wordExists := make(map[string]struct{})

	if len(pattern) != len(sArr) {
		return false
	}

	for i, v := range pattern {
		if word, ok := dict[v]; !ok {
			if _, ok := wordExists[sArr[i]]; ok {
				return false
			}
			dict[v] = sArr[i]
		} else if word != sArr[i] {
			return false
		}
		wordExists[sArr[i]] = struct{}{}
	}
	return true
}
