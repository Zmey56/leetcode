package canConstruct

//Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters
//from magazine and false otherwise.
//
//Each letter in magazine can only be used once in ransomNote.

func canConstruct(ransomNote string, magazine string) bool {
	result := make(map[rune]int)

	for _, v := range ransomNote {
		result[v]++
	}

	for _, v := range magazine {
		if result[v] > 0 {
			result[v]--
		}
	}

	for _, v := range result {
		if v > 0 {
			return false
		}
	}

	return true
}
