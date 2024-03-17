package canConstruct

import "testing"

// Test for canConstruct
func TestCanConstruct(t *testing.T) {

	test := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, test := range test {
		result := canConstruct(test.ransomNote, test.magazine)
		if result != test.expected {
			t.Errorf("Test failed: expected %v, got %v", test.expected, result)
		}
	}
}
