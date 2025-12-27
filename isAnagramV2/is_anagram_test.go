package isAnagramV2

import "testing"

func TestIsAnagram(t *testing.T) {
	test := []struct {
		s      string
		t      string
		output bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ab", "ba", true},
		{"a", "a", true},
	}

	for _, testCase := range test {
		if isAnagram(testCase.s, testCase.t) != testCase.output {
			t.Errorf("Expected %v but got %v", testCase.output, isAnagram(testCase.s, testCase.t))
		}
	}
}
