package uncommonFromSentences

import "strings"

//A sentence is a string of single-space separated words where each word consists only of lowercase letters.
//
//A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
//
//Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.

func uncommonFromSentences(s1 string, s2 string) []string {
	mapWords := make(map[string]int)
	for _, word := range append(splitWords(s1), splitWords(s2)...) {
		mapWords[word]++
	}

	var result []string
	for word, count := range mapWords {
		if count == 1 {
			result = append(result, word)
		}
	}

	return result
}

func splitWords(s string) []string {
	return strings.Split(s, " ")
}

func uncommonFromSentencesVer2(s1 string, s2 string) []string {
	mapWords := make(map[string]int)

	// Split and iterate over words from s1
	for _, word := range splitWords(s1) {
		mapWords[word]++
	}

	// Split and iterate over words from s2
	for _, word := range splitWords(s2) {
		mapWords[word]++
	}

	var result []string
	for word, count := range mapWords {
		if count == 1 {
			result = append(result, word)
		}
	}

	return result
}
