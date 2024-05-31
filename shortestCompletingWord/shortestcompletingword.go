package shortestCompletingWord

//Given a string licensePlate and an array of strings words, find the shortest completing word in words.
//
//A completing word is a word that contains all the letters in licensePlate. Ignore numbers and spaces in licensePlate,
//and treat letters as case insensitive. If a letter appears more than once in licensePlate, then it must appear in the word the same number of times or more.
//
//For example, if licensePlate = "aBc 12c", then it contains letters 'a', 'b' (ignoring case), and 'c' twice.
//Possible completing words are "abccdef", "caaacab", and "cbca".
//
//Return the shortest completing word in words. It is guaranteed an answer exists. If there are multiple shortest completing words,
//return the first one that occurs in words.

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlateMap := make(map[rune]int)
	for _, char := range licensePlate {
		if char >= 'a' && char <= 'z' {
			licensePlateMap[char]++
		} else if char >= 'A' && char <= 'Z' {
			licensePlateMap[char+32]++
		}
	}

	wordMap := make(map[rune]int)
	shortestWord := ""
	for _, word := range words {
		for _, char := range word {
			wordMap[char]++
		}
		isCompleting := true
		for char, count := range licensePlateMap {
			if wordMap[char] < count {
				isCompleting = false
				break
			}
		}
		if isCompleting {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
		}
		wordMap = make(map[rune]int)
	}

	return shortestWord

}
