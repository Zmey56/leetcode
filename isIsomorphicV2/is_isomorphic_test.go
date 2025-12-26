package isIsomorphicV2

import "testing"

func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		s      string
		t      string
		result bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"ab", "aa", false},
		{"", "", true},
	}

	for _, testCase := range testCases {
		result := isIsomorphic(testCase.s, testCase.t)
		if result != testCase.result {
			t.Errorf("For s = %s and t = %s, expected %v, but got %v", testCase.s, testCase.t, testCase.result, result)
		}
	}

}
