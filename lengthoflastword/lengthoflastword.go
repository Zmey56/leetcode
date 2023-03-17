//58. Length of Last Word
//
//Given a string s consisting of words and spaces, return the length of the last word in the string.
//
//A word is a maximal
//substring
//consisting of non-space characters only.
//
//
//
//Example 1:
//
//Input: s = "Hello World"
//Output: 5
//Explanation: The last word is "World" with length 5.
//Example 2:
//
//Input: s = "   fly me   to   the moon  "
//Output: 4
//Explanation: The last word is "moon" with length 4.
//Example 3:
//
//Input: s = "luffy is still joyboy"
//Output: 6
//Explanation: The last word is "joyboy" with length 6.
//
//
//Constraints:
//
//1 <= s.length <= 104
//s consists of only English letters and spaces ' '.
//There will be at least one word in s.

package main

import (
	"log"
	"strings"
)

func lengthOfLastWord(s string) int {
	// Trim spaces from the end of the string
	s = strings.TrimRight(s, " ")

	// Find the last space in the string
	lastSpaceIndex := strings.LastIndex(s, " ")

	// If there is no space, the entire string is the last word
	if lastSpaceIndex == -1 {
		return len(s)
	}

	// Otherwise, return the length of the last word
	return len(s) - lastSpaceIndex - 1
}

func main() {
	s := "Hello World"
	log.Println(lengthOfLastWord(s))

	s2 := "   fly me   to   the moon  "
	log.Println(lengthOfLastWord(s2))

	s3 := "luffy is still joyboy"
	log.Println(lengthOfLastWord(s3))
}
