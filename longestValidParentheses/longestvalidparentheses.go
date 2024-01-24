package longestValidParentheses

import "log"

//Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed)
// parentheses substring.

func longestValidParentheses(s string) int {
	var res int
	stack := []int{-1}
	log.Println("stack ", stack)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
			log.Println("c == '(' ", stack)
		} else {
			stack = stack[:len(stack)-1]
			log.Println("c == ')' ", stack)
			if len(stack) == 0 {
				stack = append(stack, i)
				log.Println("len(stack) == 0 ", stack)
			} else {
				res = max(res, i-stack[len(stack)-1])
				log.Println("res = max(res, i-stack[len(stack)-1]) ", res)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
