package preorderTraversal

import "testing"

// Test for preorderTraversal function
func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		example  string
		input    *TreeNode
		expected []int
	}{
		{
			example:  "Test 1",
			input:    &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: nil}},
			expected: []int{1, 2, 3},
		},
		{
			example:  "Test 2",
			input:    &TreeNode{Val: 1, Left: nil, Right: nil},
			expected: []int{1},
		},
		{
			example:  "Test 3",
			input:    nil,
			expected: []int{},
		},
	}

	for _, test := range tests {
		result := preorderTraversal(test.input)
		if !equal(result, test.expected) {
			t.Errorf("Test %s failed, expected %v, got %v", test.example, test.expected, result)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
