package main

func isSubsequence(s string, t string) bool {
	sIdx := 0 // Index for iterating through string s
	tIdx := 0 // Index for iterating through string t

	for tIdx < len(t) && sIdx < len(s) {
		if s[sIdx] == t[tIdx] {
			sIdx++
		}
		tIdx++
	}

	return sIdx == len(s)
}
