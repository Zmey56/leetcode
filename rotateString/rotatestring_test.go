package rotateString

import "testing"

func Test_rotateString(t *testing.T) {
	tests := []struct {
		name string
		A    string
		B    string
		want bool
	}{
		{"Example 1", "abcde", "cdeab", true},
		{"Example 2", "abcde", "abced", false},
		{"Example 3", "abcde", "abcde", true},
		{"Example 4", "abcde", "eabcd", true},
		{"Example 5", "abcde", "eabdc", false},
		{"Example 6", "abcde", "abcdf", false},
		{"Example 7", "abcde", "fabcde", false},
		{"Example 8", "abcde", "efabcd", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.A, tt.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
