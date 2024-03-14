package rangeSumQuery

import "testing"

//Test for function NumArray

func TestNumArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		left     int
		right    int
		expected int
	}{
		{
			name:     "Test Case 1",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			left:     0,
			right:    2,
			expected: 1,
		},
		{
			name:     "Test Case 2",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			left:     2,
			right:    5,
			expected: -1,
		},
		{
			name:     "Test Case 3",
			nums:     []int{-2, 0, 3, -5, 2, -1},
			left:     0,
			right:    5,
			expected: -3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			obj := Constructor(test.nums)
			actual := obj.SumRange(test.left, test.right)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}
