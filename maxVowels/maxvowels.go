package main

import "strings"

func maxVowels(s string, k int) int {
	vowels := "aeiou"
	maxCount := 0
	currentCount := 0

	for i := 0; i < len(s); i++ {
		if i >= k {
			if strings.Contains(vowels, string(s[i-k])) {
				currentCount--
			}
		}

		if strings.Contains(vowels, string(s[i])) {
			currentCount++
		}

		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
}
