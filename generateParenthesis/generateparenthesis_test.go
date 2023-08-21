package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n      int
		output []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{2, []string{"(())", "()()"}},
		{1, []string{"()"}},
	}

	for _, test := range tests {
		result := generateParenthesis(test.n)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("For n=%d, expected %v but got %v", test.n, test.output, result)
		}
	}

}
