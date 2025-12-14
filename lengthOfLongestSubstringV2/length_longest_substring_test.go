package lengthOfLongestSubstringV2

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
	}

	for _, tt := range tests {
		b.Run("BenchmarkLengthOfLongestSubstring", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lengthOfLongestSubstring(tt.input)
			}
		})

		b.Run("BenchmarkLengthOfLongestSubstring", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lengthOfLongestSubstringV2(tt.input)
			}
		})
	}
}
