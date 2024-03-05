package reverseBits

import "testing"

//Test for function reverseBits

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{
			name:     "Test Case 1",
			input:    43261596,
			expected: 964176192,
		},
		{
			name:     "Test Case 2",
			input:    4294967293,
			expected: 3221225471,
		},
		{
			name:     "Test Case 3",
			input:    0,
			expected: 0,
		},
		{
			name:     "Test Case 4",
			input:    1,
			expected: 2147483648,
		},
		{
			name:     "Test Case 5",
			input:    2147483648,
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := reverseBits(test.input)
			if got != test.expected {
				t.Errorf("Got %v, expected %v", got, test.expected)
			}
		})
	}
}
