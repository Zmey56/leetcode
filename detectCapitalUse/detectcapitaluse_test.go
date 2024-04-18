package detectCapitalUse

import "testing"

// Test for function detectCapitalUse
func Test_detectCapitalUse(t *testing.T) {
	tests := []struct {
		word   string
		result bool
	}{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
		{"mL", false},
	}

	for _, tt := range tests {
		if got := detectCapitalUse(tt.word); got != tt.result {
			t.Errorf("detectCapitalUse(%v) = %v, want %v", tt.word, got, tt.result)
		}
	}
}
