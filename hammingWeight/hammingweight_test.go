package hammingWeight

import "testing"

// Test for function hammingWeight
func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected int
	}{
		{
			name:     "Test Case 1",
			input:    00000000000000000000000000001011,
			expected: 3,
		},
		{
			name:     "Test Case 2",
			input:    00000000000000000000000010000000,
			expected: 1,
		},
		{
			name:     "Test Case 3",
			input:    11111111111111111111111111111101,
			expected: 31,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := hammingWeight(test.input)
			if got != test.expected {
				t.Errorf("Got %v, expected %v", got, test.expected)
			}
		})
	}
}
