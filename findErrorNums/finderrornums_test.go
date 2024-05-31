package findErrorNums

import "testing"

// Test for function findErrorNums
func Test_findErrorNums(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test case 1",
			nums:     []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		{
			name:     "Test case 2",
			nums:     []int{1, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Test case 3",
			nums:     []int{3, 2, 3, 4, 6, 5},
			expected: []int{3, 1},
		},
		{
			name:     "Test case 4",
			nums:     []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9},
			expected: []int{2, 10},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findErrorNums(test.nums)
			if !equal(result, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, result)
			}
		})
	}

}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
