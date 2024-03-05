package isHappy

import "testing"

//Test for function isHappy

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Test Case 1",
			input:    19,
			expected: true,
		},
		{
			name:     "Test Case 2",
			input:    2,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isHappy(test.input)
			if got != test.expected {
				t.Errorf("Got %v, expected %v", got, test.expected)
			}
		})
	}
}
