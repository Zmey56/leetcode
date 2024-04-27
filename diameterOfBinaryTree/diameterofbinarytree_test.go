package diameterOfBinaryTree

import "testing"

// Test for function diameterOfBinaryTree

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result int
	}{
		{&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}, 3},
		{&TreeNode{1, nil, nil}, 0},
		{&TreeNode{1, &TreeNode{2, nil, nil}, nil}, 1},
		{&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}, 2},
	}

	for _, tt := range tests {
		if got := diameterOfBinaryTree(tt.root); got != tt.result {
			t.Errorf("diameterOfBinaryTree(%v) = %v, want %v", tt.root, got, tt.result)
		}
	}
}
