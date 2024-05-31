package buddyStrings

//Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.
//
//Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].
//
//For example, swapping at indices 0 and 2 in "abcd" results in "cbad".

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		// check if there are duplicate characters in s
		// if there are, we can swap them
		// if there are none, we can't swap any characters
		seen := make(map[rune]bool)
		for _, c := range s {
			if seen[c] {
				return true
			}
			seen[c] = true
		}
		return false
	}

	// find the first pair of characters that are different
	// in s and goal
	first := -1
	second := -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	// if we found a pair of characters that are different
	// in s and goal, check if swapping them results in goal
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
