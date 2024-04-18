package getMinimumDifference

import "testing"

// Test for function getMinimumDifference
func Test_getMinimumDifference(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result int
	}{
		{&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}}, 1},
		{&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 6, Left: nil, Right: nil}}, 1},
		{&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}, 2},
	}

	for _, tt := range tests {
		if got := getMinimumDifference(tt.root); got != tt.result {
			t.Errorf("getMinimumDifference(%v) = %v, want %v", tt.root, got, tt.result)
		}
	}
}
