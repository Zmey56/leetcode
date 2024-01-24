package longestValidParentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	test := []struct {
		s    string
		want int
	}{
		{
			s:    "(()",
			want: 2,
		},
		{
			s:    ")()())",
			want: 4,
		},
		{
			s:    "",
			want: 0,
		},
		{
			s:    "(()())",
			want: 6,
		},
		{
			s:    "(()())(",
			want: 6,
		},
		{
			s:    "(()())()",
			want: 8,
		},
		{
			s:    "(()())()(",
			want: 8,
		},
		{
			s:    "(()())()()",
			want: 10,
		},
		{
			s:    "(()())()()(",
			want: 10,
		},
	}
	for _, test := range test {
		got := longestValidParentheses(test.s)
		if got != test.want {
			t.Errorf("longestValidParentheses(%s) = %d, want %d", test.s, got, test.want)
		}
	}
}
