package licenseKeyFormatting

import (
	"strings"
)

//You are given a license key represented as a string s that consists of only alphanumeric characters and dashes.
//The string is separated into n + 1 groups by n dashes. You are also given an integer k.
//
//We want to reformat the string s such that each group contains exactly k characters, except for the first group,
//which could be shorter than k but still must contain at least one character. Furthermore,
//there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.
//
//Return the reformatted license key

func licenseKeyFormatting(s string, k int) string {
	s = strings.ToUpper(strings.Replace(s, "-", "", -1))
	l := len(s)
	runes := []rune{}

	for i, ru := range s {
		runes = append(runes, ru)
		if i < l-1 && (l-i-1)%k == 0 {
			runes = append(runes, '-')
		}
	}

	return string(runes)
}
