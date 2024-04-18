package findWords

import "strings"

//Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.
//
//In the American keyboard:
//
//the first row consists of the characters "qwertyuiop",
//the second row consists of the characters "asdfghjkl", and
//the third row consists of the characters "zxcvbnm".

func findWords(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	result := []string{}
	firstLine := "qwertyuiop"
	secondLine := "asdfghjkl"
	thirdLine := "zxcvbnm"

	first := []rune(firstLine)
	second := []rune(secondLine)
	third := []rune(thirdLine)

	for _, word := range words {
		wordLower := strings.ToLower(word)
		count := 0
		firstLineCount := 0
		secondLineCount := 0
		thirdLineCount := 0
		for _, letter := range wordLower {
			count++
			if contains(first, letter) {
				firstLineCount++
			}
			if contains(second, letter) {
				secondLineCount++
			}
			if contains(third, letter) {
				thirdLineCount++
			}
		}
		if firstLineCount == count || secondLineCount == count || thirdLineCount == count {
			result = append(result, word)
		}
	}

	return result
}

func contains(arr []rune, letter rune) bool {
	for _, a := range arr {
		if a == letter {
			return true
		}
	}
	return false
}
