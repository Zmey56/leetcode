package isIsomorphic

//Given two strings s and t, determine if they are isomorphic.
//
//Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//
//All occurrences of a character must be replaced with another character while preserving the order of characters.
//No two characters may map to the same character, but a character may map to itself.

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] == 0 {
			m1[s[i]] = t[i]
		} else {
			if m1[s[i]] != t[i] {
				return false
			}
		}
		if m2[t[i]] == 0 {
			m2[t[i]] = s[i]
		} else {
			if m2[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}
