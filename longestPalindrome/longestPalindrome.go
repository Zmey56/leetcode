package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	tmp := make([]uint8, len(s))
	result := []uint8{}
	length := utf8.RuneCountInString(s)
	if length == 1 {
		return s
	} else if length > 1 {
		for i := 0; i < len(s); i++ {
			for j := len(s) - 1; j > i; j-- {
				if s[i] == s[j] {
					for start, end := i, j; end >= start; start, end = start+1, end-1 {
						if start <= end {
							//fmt.Println(string(s[start]), string(s[end]), start, end)
							if s[start] == s[end] {
								tmp[start] = s[start]
								tmp[end] = s[end]
								//fmt.Println("TMP", string(tmp[start]), string(tmp[end]), start, end, lenArrayWithZero(tmp))
								//if start == end {
								//	result = tmp
								//	//fmt.Println("R", tmp, len(result), result)
								//	if lenArrayWithZero(result) < lenArrayWithZero(tmp) {
								//		result = tmp
								//		fmt.Println("R2", result)
								//	}
								//}
							} else {
								fmt.Println(string(tmp))
								tmp = tmp[:0]
								break
							}
						}

					}
					//fmt.Println("OK", tmp, lenArrayWithZero(result), lenArrayWithZero(tmp))
					if lenArrayWithZero(result) < lenArrayWithZero(tmp) {
						result = tmp
					}
					tmp = make([]uint8, len(s))
				}
			}
		}
		if lenArrayWithZero(result) < 1 {
			return string(s[0])
		}
	}

	return (strings.Trim(string(result), "\u0000"))
}

func lenArrayWithZero(s []uint8) int {
	nonZeroCount := 0

	for _, value := range s {
		if value != 0 {
			nonZeroCount++
		}
	}
	return nonZeroCount
}
