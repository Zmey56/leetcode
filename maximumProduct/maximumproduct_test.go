package maximumProduct

import "testing"

//Test for function maximumProduct

func Test_maximumProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test case 1",
			nums:     []int{1, 2, 3, 4},
			expected: 24,
		},
		{
			name:     "Test case 2",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Test case 3",
			nums:     []int{-1, -2, -3},
			expected: -6,
		},
		{
			name:     "Test case 4",
			nums:     []int{-1, -2, 1},
			expected: 2,
		},
		{
			name:     "Test case 5",
			nums:     []int{-1, -2, 1, 2, 3},
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maximumProduct(test.nums)
			if result != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, result)
			}
		})
	}

}
