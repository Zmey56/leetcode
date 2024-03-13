package canWinNim

import "testing"

// Test for function canWinNim

func TestCanWinNim(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Test Case 1",
			n:        4,
			expected: false,
		},
		{
			name:     "Test Case 2",
			n:        1,
			expected: true,
		},
		{
			name:     "Test Case 3",
			n:        2,
			expected: true,
		},
		{
			name:     "Test Case 4",
			n:        7,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := canWinNim(test.n)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}

}
