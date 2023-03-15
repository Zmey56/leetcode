//28. Find the Index of the First Occurrence in a String
//
//Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
//
//Example 1:
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//Example 2:
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//
//Constraints:
//
//1 <= haystack.length, needle.length <= 104
//haystack and needle consist of only lowercase English characters.

package main

import "log"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i < n-m+1; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	log.Println(strStr(haystack, needle))

	haystack1 := "leetcode"
	needle1 := "leeto"
	log.Println(strStr(haystack1, needle1))
}
