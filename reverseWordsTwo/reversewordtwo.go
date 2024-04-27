package reverseWordsTwo

import "strings"

//Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

func reverseWords(s string) string {
	splitWors := strings.Split(s, " ")
	result := ""

	for _, word := range splitWors {
		result += reverseWord(word) + " "
	}

	return strings.Trim(result, " ")
}

func reverseWord(word string) string {
	result := ""
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}
