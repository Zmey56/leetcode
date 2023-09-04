package main

import (
	"fmt"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := []string{}
	numStack := []int{}
	currentNum := 0
	currentStr := ""

	for _, char := range s {
		if unicode.IsDigit(char) {
			currentNum = currentNum*10 + int(char-'0')
		} else if char == '[' {
			numStack = append(numStack, currentNum)
			stack = append(stack, currentStr)

			currentNum = 0
			currentStr = ""
		} else if char == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			currentStr = prevStr + strings.Repeat(currentStr, num)
		} else {
			currentStr += string(char)
		}
	}

	return currentStr
}

func main() {
	// Test cases
	s1 := "3[a]2[bc]"
	fmt.Println(decodeString(s1))

	s2 := "3[a2[c]]"
	fmt.Println(decodeString(s2))

	s3 := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s3))

}
