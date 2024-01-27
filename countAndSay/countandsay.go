package countAndSay

import (
	"strconv"
)

// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
//
//countAndSay(1) = "1"
//countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into
//a different digit string.
//To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring
//contains exactly one unique digit. Then for each substring, say the number of digits, then say the digit.
//Finally, concatenate every said digit.
//
//For example, the saying and conversion for digit string "3322251":
//
//
//Given a positive integer n, return the nth term of the count-and-say sequence.

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return say(countAndSay(n - 1))
}

func say(s string) string {
	var result string
	var count int
	for i := 0; i < len(s); i++ {
		count++
		if i == len(s)-1 || s[i] != s[i+1] {
			result += strconv.Itoa(count) + string(s[i])
			count = 0
		}
	}
	return result
}
