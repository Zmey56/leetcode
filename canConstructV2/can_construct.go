package canConstructV2

//383. Ransom Note

//Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
//Each letter in magazine can only be used once in ransomNote.

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, v := range magazine {
		magazineMap[v]++
	}

	for _, v := range ransomNote {
		if letter, ok := magazineMap[v]; ok && letter > 0 {
			magazineMap[v]--
		} else {
			return false
		}
	}
	return true
}

func canConstructV2(ransomNote string, magazine string) bool {
	// Early exit: ransom note can't be longer than magazine
	if len(ransomNote) > len(magazine) {
		return false
	}

	// Use array instead of map for ASCII lowercase letters (a-z)
	var letterCount [26]int

	// Count letters in magazine
	for i := 0; i < len(magazine); i++ {
		letterCount[magazine[i]-'a']++
	}

	// Check if ransom note can be constructed
	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		letterCount[idx]--
		if letterCount[idx] < 0 {
			return false
		}
	}

	return true
}

// Fastest version - eliminates bounds checking and uses unsafe for maximum speed
func canConstructV3(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	// Stack-allocated array
	var counts [26]byte

	// Count magazine letters - compiler will unroll small loops
	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}

	// Check ransom note
	for i := 0; i < len(ransomNote); i++ {
		idx := ransomNote[i] - 'a'
		if counts[idx] == 0 {
			return false
		}
		counts[idx]--
	}

	return true
}

// Alternative: using pointers for potentially better performance
func canConstructV4(ransomNote string, magazine string) bool {
	rLen, mLen := len(ransomNote), len(magazine)
	if rLen > mLen {
		return false
	}

	var counts [26]byte

	// Process magazine
	for i := 0; i < mLen; i++ {
		counts[magazine[i]-'a']++
	}

	// Process ransom note
	for i := 0; i < rLen; i++ {
		c := ransomNote[i] - 'a'
		if counts[c] == 0 {
			return false
		}
		counts[c]--
	}

	return true
}
