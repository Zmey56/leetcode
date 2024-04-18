package detectCapitalUse

import "unicode"

//We define the usage of capitals in a word to be right when one of the following cases holds:
//
//All letters in this word are capitals, like "USA".
//All letters in this word are not capitals, like "leetcode".
//Only the first letter in this word is capital, like "Google".
//Given a string word, return true if the usage of capitals in it is right.

func detectCapitalUse(word string) bool {
	firstCapital := false
	for i, c := range word {
		if i == 0 && unicode.IsUpper(c) {
			firstCapital = true
		}
		if i == 1 {
			if firstCapital && unicode.IsUpper(c) {
				firstCapital = true
			} else if firstCapital && unicode.IsLower(c) {
				firstCapital = false
			} else if !firstCapital && unicode.IsUpper(c) {
				return false
			}
		}
		if i > 1 {
			if firstCapital && unicode.IsUpper(c) {
				continue
			}
			if !firstCapital && unicode.IsLower(c) {
				continue
			}
			return false
		}

	}
	return true

}
