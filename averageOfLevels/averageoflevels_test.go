package averageOfLevels

import "testing"

// Test for function averageOfLevels
func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []float64
	}{
		{
			name:     "Test case 1",
			root:     &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
			expected: []float64{3, 14.5, 11},
		},
		{
			name:     "Test case 2",
			root:     &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, &TreeNode{5, nil, nil}, nil}},
			expected: []float64{1, 2.5, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := averageOfLevels(test.root)
			if !equal(result, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, result)
			}
		})
	}
}

func equal(a, b []float64) bool {
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
