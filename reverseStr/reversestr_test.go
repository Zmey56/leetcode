package reverseStr

import "testing"

// Test for function reverseStr
func Test_reverseStr(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		result string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
		{"abcdefg", 3, "cbadefg"},
		{"abcdefg", 4, "dcbaefg"},
	}

	for _, tt := range tests {
		if got := reverseStr(tt.s, tt.k); got != tt.result {
			t.Errorf("reverseStr(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.result)
		}
	}
}
