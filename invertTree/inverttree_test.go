package invertTree

import "testing"

//Test for invertTree

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	inverted := inverTree(root)
	if inverted.Val != 4 {
		t.Errorf("Expected 4 but got %d", inverted.Val)
	}
	if inverted.Left.Val != 7 {
		t.Errorf("Expected 7 but got %d", inverted.Left.Val)
	}
	if inverted.Right.Val != 2 {
		t.Errorf("Expected 2 but got %d", inverted.Right.Val)
	}
	if inverted.Left.Left.Val != 9 {
		t.Errorf("Expected 9 but got %d", inverted.Left.Left.Val)
	}
	if inverted.Left.Right.Val != 6 {
		t.Errorf("Expected 6 but got %d", inverted.Left.Right.Val)
	}
	if inverted.Right.Left.Val != 3 {
		t.Errorf("Expected 3 but got %d", inverted.Right.Left.Val)
	}
	if inverted.Right.Right.Val != 1 {
		t.Errorf("Expected 1 but got %d", inverted.Right.Right.Val)
	}

}
