package toGoatLatin

import "strings"

//You are given a string sentence that consist of words separated by spaces. Each word consists of lowercase and uppercase letters only.
//
//We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.) The rules of Goat Latin are as follows:
//
//If a word begins with a vowel ('a', 'e', 'i', 'o', or 'u'), append "ma" to the end of the word.
//For example, the word "apple" becomes "applema".
//If a word begins with a consonant (i.e., not a vowel), remove the first letter and append it to the end, then add "ma".
//For example, the word "goat" becomes "oatgma".
//Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
//For example, the first word gets "a" added to the end, the second word gets "aa" added to the end, and so on.
//Return the final sentence representing the conversion from sentence to Goat Latin.

func toGoatLatin(sentence string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	ma := "ma"

	result := ""

	separetedWords := strings.Split(sentence, " ")
	for i, word := range separetedWords {
		if _, ok := vowels[word[0]]; ok {
			word += ma
		} else {
			word = word[1:] + string(word[0]) + ma
		}
		word += strings.Repeat("a", i+1)
		result += word + " "

	}

	return result[:len(result)-1]
}
