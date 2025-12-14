package lengthOfLongestSubstringV2

//Given a string s, find the length of the longest substring without duplicate characters.
//
//
//
//Example 1:
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
//Example 2:
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//Example 3:
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	mapLetter := make(map[rune]struct{})
	maxLongest := 0
	count := 0

	for i := 0; i < len(s); i++ {
		count++
		mapLetter[rune(s[i])] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, exists := mapLetter[rune(s[j])]; !exists {
				count++
				mapLetter[rune(s[j])] = struct{}{}
			} else {
				if count > maxLongest {
					maxLongest = count
				}
				count = 0
				mapLetter = make(map[rune]struct{})
				break
			}

			if j == len(s)-1 {
				if count > maxLongest {
					maxLongest = count
				}
				count = 0
			}

		}
	}
	if count > maxLongest {
		maxLongest = count
	}

	return maxLongest
}

func lengthOfLongestSubstringV2(s string) int {
	charIndex := make(map[rune]int)
	maxLen := 0
	left := 0

	for right, char := range s {
		// If character exists and is in current window
		if lastIdx, exists := charIndex[char]; exists && lastIdx >= left {
			left = lastIdx + 1
		}

		charIndex[char] = right

		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
