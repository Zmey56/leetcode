package moveZeroesTwo

import "testing"

//Test for function moveZeroes

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test Case 1",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Test Case 2",
			input:    []int{0},
			expected: []int{0},
		},
		{
			name:     "Test Case 3",
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Test Case 4",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test Case 5",
			input:    []int{0, 0, 0, 0, 1},
			expected: []int{1, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moveZeroes(test.input)
			for i := 0; i < len(test.input); i++ {
				if test.input[i] != test.expected[i] {
					t.Errorf("Expected %v but got %v", test.expected, test.input)
				}
			}
		})
	}
}
