// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	//if len(strs) == 0 {
	//	return ""
	//}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			log.Println(strs[i], " - ", prefix)
			if len(prefix) == 0 {
				fmt.Println("NONE")
			}
		}
	}

	fmt.Println(prefix)
}
