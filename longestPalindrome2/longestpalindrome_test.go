package longestPalindrome2

import "testing"

// Test for function longestPalindrome
func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "abccccdd",
			want: 7,
		},
		{
			s:    "a",
			want: 1,
		},
		{
			s:    "bb",
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := longestPalindrome(tt.s); got != tt.want {
			t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
		}
	}

}
