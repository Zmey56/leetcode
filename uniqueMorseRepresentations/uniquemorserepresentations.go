package uniqueMorseRepresentations

import (
	"runtime"
	"sync"
)

// International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows:
//
//'a' maps to ".-",
//'b' maps to "-...",
//'c' maps to "-.-.", and so on.
//For convenience, the full table for the 26 letters of the English alphabet is given below:
//
//[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
//Given an array of strings words where each word can be written as a concatenation of the Morse code of each letter.
//
//For example, "cab" can be written as "-.-..--...", which is the concatenation of "-.-.", ".-", and "-...". We will call such a concatenation the transformation of a word.
//Return the number of different transformations among all words we have.

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	numGoroutines := runtime.NumCPU()
	maps := make([]map[string]bool, numGoroutines)

	var wg sync.WaitGroup
	partSize := len(words) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if i == numGoroutines-1 {
			end = len(words)
		}

		wg.Add(1)
		go func(part []string, i int) {
			defer wg.Done()
			maps[i] = make(map[string]bool) // Initialize the map
			for _, word := range part {
				morse := ""
				for _, c := range word {
					morse += morseCodes[c-'a']
				}
				maps[i][morse] = true
			}
		}(words[start:end], i)
	}

	wg.Wait()

	morseMap := make(map[string]bool)
	for _, m := range maps {
		for k := range m {
			morseMap[k] = true
		}
	}

	return len(morseMap)
}
