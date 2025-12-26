package wordPatternV2

import "testing"

func TestWordPattern(t *testing.T) {
	testCases := []struct {
		pattern string
		str     string
		want    bool
	}{
		//{"abba", "dog cat cat dog", true},
		//{"abba", "dog cat cat fish", false},
		//{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, testCase := range testCases {
		result := wordPattern(testCase.pattern, testCase.str)
		if result != testCase.want {
			t.Errorf("For pattern %v and str %v, expected %v, but got %v", testCase.pattern, testCase.str, testCase.want, result)
		}
	}
}
