package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	result := []string{}

	if len(word1) == len(word2) {
		for i, j := range word1 {
			result = append(result, string(j))
			result = append(result, string(word2[i]))
		}
	} else if len(word1) > len(word2) {
		for i, j := range word1 {
			if i < len(word2) {
				result = append(result, string(j))
				result = append(result, string(word2[i]))
			} else {
				result = append(result, string(j))
			}
		}
	} else {
		for i, j := range word2 {
			if i < len(word1) {
				result = append(result, string(word1[i]))
				result = append(result, string(j))
			} else {
				result = append(result, string(j))
			}
		}
	}

	return strings.Join(result, "")
}
