package main

// 438. Find All Anagrams in a String

// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
//
//	You may return the answer in any order.
func findAnagrams(s string, p string) []int {
	var result []int

	if len(s) < len(p) {
		return result
	}

	var pCount, sCount [256]int

	pLen := len(p)
	for i := 0; i < pLen; i++ {
		pCount[p[i]]++
		sCount[s[i]]++
	}

	if pCount == sCount {
		result = append(result, 0)
	}

	for i := pLen; i < len(s); i++ {
		sCount[s[i]]++
		sCount[s[i-pLen]]--

		if pCount == sCount {
			result = append(result, i-pLen+1)
		}
	}

	return result

}
