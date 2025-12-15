package minWindow

//Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every
//character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
//
//The testcases will be generated such that the answer is unique.

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(t)
	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		if need[c] > 0 {
			required--
		}
		need[c]--

		for required == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			lc := s[left]
			need[lc]++
			if need[lc] > 0 {
				required++
			}
			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
