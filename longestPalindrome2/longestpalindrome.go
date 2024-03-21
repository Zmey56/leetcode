package longestPalindrome2

import "log"

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	log.Println("MAP", m)

	var res int
	var odd bool
	for _, v := range m {
		log.Println("V", v, "RES", res, "ODD", odd)
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
			odd = true
		}
	}
	if odd {
		res++
	}
	return res
}
