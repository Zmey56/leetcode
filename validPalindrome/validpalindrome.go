package validPalindrome

//Given a string s, return true if the s can be palindrome after deleting at most one character from it.

func validPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			if i == j-1 {
				return true
			}
			if s[i] == s[j-1] {
				j = j - 1
			} else if s[i+1] == s[j] {
				i = i + 1
			} else {
				return false
			}
		}
		i++
		j--
	}
	return true
}
