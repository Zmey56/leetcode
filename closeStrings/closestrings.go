package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, ch := range word1 {
		freq1[ch]++
	}

	for _, ch := range word2 {
		freq2[ch]++
	}

	for ch := range freq1 {
		if _, exists := freq2[ch]; !exists {
			return false
		}
	}

	freqCount1 := make(map[int]int)
	freqCount2 := make(map[int]int)

	for _, freq := range freq1 {
		freqCount1[freq]++
	}

	for _, freq := range freq2 {
		freqCount2[freq]++
	}

	for freq, count := range freqCount1 {
		if freqCount2[freq] != count {
			return false
		}
	}

	return true
}
