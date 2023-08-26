package main

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	strRunes := []rune(s)
	left, right := 0, len(strRunes)-1

	for left < right {
		for left < right && !strings.ContainsRune(vowels, strRunes[left]) {
			left++
		}
		for left < right && !strings.ContainsRune(vowels, strRunes[right]) {
			right--
		}

		// Swap the vowels
		strRunes[left], strRunes[right] = strRunes[right], strRunes[left]
		left++
		right--
	}

	return string(strRunes)
}
