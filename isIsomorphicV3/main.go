package main

import (
	"fmt"
)

//205. Isomorphic Strings

// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving
//
//	the order of characters. No two characters may map to the same character, but a character may map to itself.
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[byte]byte)
	used := make(map[byte]bool)
	// runeS := []rune(s)
	// runeT := []rune(t)

	for i := 0; i < len(s); i++ {
		fmt.Printf("runeS, %s\n", string(l))
		fmt.Println("MAP ", letters)
		if m, ok := letters[s[i]]; ok {
			if m != t[i] {
				return false
			}
			continue
		}
		if used[t[i]] {
			return false
		}

		letters[s[i]] = t[i]
		used[t[i]] = true
	}
	return true
}
