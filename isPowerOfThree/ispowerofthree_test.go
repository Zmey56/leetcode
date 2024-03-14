package isPowerOfThree

import "testing"

//Test for function isPowerOfThree

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Test Case 1",
			n:        27,
			expected: true,
		},
		{
			name:     "Test Case 2",
			n:        0,
			expected: false,
		},
		{
			name:     "Test Case 3",
			n:        9,
			expected: true,
		},
		{
			name:     "Test Case 4",
			n:        45,
			expected: false,
		},
		{
			name:     "Test Case 5",
			n:        1,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isPowerOfThree(test.n)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}
