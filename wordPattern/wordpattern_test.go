package wordPattern

import "testing"

//Test for function wordPattern

func TestWordPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		s        string
		expected bool
	}{
		{
			name:     "Test Case 1",
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		{
			name:     "Test Case 2",
			pattern:  "abba",
			s:        "dog cat cat fish",
			expected: false,
		},
		{
			name:     "Test Case 3",
			pattern:  "aaaa",
			s:        "dog cat cat dog",
			expected: false,
		},
		{
			name:     "Test Case 4",
			pattern:  "abba",
			s:        "dog dog dog dog",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := wordPattern(test.pattern, test.s)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}

}
