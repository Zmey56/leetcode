package main

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)      // Split the input string into words
	reverseSlice(words)             // Reverse the slice of words
	return strings.Join(words, " ") // Join the words with a single space
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
