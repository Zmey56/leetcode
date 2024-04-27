package isSubtree

import "testing"

// Test for function sSubtree

func Test_sSubtree(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}

	result := sSubtree(root, subRoot)
	if result != true {
		t.Errorf("sSubtree(root, subRoot) = %v; want true", result)
	}
}
