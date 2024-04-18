package findLUSlength

import "testing"

// Test for function findLUSlength
func Test_findLUSlength(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		result int
	}{
		{"aba", "cdc", 3},
		{"aaa", "aaa", -1},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
		{"aaa", "aaa", -1},
	}

	for _, tt := range tests {
		if got := findLUSlength(tt.a, tt.b); got != tt.result {
			t.Errorf("findLUSlength(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.result)
		}
	}
}
