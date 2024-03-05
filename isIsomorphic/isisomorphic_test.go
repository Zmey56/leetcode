package isIsomorphic

import "testing"

// Test cases for isIsomorphic function

func TestIsIsomorphic(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
	}

	for _, test := range tests {
		if got := isIsomorphic(test.s, test.t); got != test.want {
			t.Errorf("isIsomorphic(%q, %q) = %v", test.s, test.t, got)
		}
	}
}
