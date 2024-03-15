package isPowerOfFour

import "testing"

// Test for function isPowerOfFour
func TestIsPowerForFour(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Test Case 1",
			n:        16,
			expected: true,
		},
		{
			name:     "Test Case 2",
			n:        5,
			expected: false,
		},
		{
			name:     "Test Case 3",
			n:        1,
			expected: true,
		},
		{
			name:     "Test Case 4",
			n:        4,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isPowerOffFour(test.n)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}

}
