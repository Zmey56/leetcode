package main

// 22. Generate Parentheses

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(current string, open, close int)

	backtrack = func(current string, open, close int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if open < n {
			backtrack(current+"(", open+1, close)
		}

		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
