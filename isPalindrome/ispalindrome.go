package isPalindrome

import (
	"strings"
	"unicode"
)

//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing
//all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters
//include letters and numbers.
//
//Given a string s, return true if it is a palindrome, or false otherwise.

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return true
	}

	result := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			result += string(unicode.ToLower(r))
		} else if unicode.IsNumber(r) {
			result += string(r)
		}
	}

	endWord := len(result) - 1
	for i := 0; i < len(result); i++ {
		if result[i] != result[endWord] {
			return false
		}
		endWord--
	}

	return true

}
