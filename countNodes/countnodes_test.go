package countNodes

import "testing"

//Test for countNodes

func TestCountNodes(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}
	result := countNodes(root)
	if result != 6 {
		t.Errorf("countNodes(root) = %d; want 6", result)
	}

}
