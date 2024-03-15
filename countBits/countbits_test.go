package countBits

import "testing"

//Test for function countBits

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Test Case 1",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "Test Case 2",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countBits(test.n)
			for i := 0; i < len(actual); i++ {
				if actual[i] != test.expected[i] {
					t.Errorf("Expected %v but got %v", test.expected, actual)
				}
			}
		})
	}

}
