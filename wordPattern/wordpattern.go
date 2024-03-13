package wordPattern

import (
	"log"
	"strings"
)

//Given a pattern and a string s, find if s follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

func wordPattern(pattern string, s string) bool {
	splitString := strings.Split(s, " ")

	if len(pattern) != len(splitString) {
		return false
	}

	patternMap := make(map[byte]string)
	wordMap := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		log.Println(patternMap, wordMap, pattern[i], splitString[i])
		if val, ok := patternMap[pattern[i]]; ok {
			if val != splitString[i] {
				return false
			}
		} else {
			patternMap[pattern[i]] = splitString[i]
		}

		if val, ok := wordMap[splitString[i]]; ok {
			if val != pattern[i] {
				return false
			}
		} else {
			wordMap[splitString[i]] = pattern[i]
		}
	}

	return true
}
