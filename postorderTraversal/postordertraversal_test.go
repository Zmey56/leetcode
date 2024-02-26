package postorderTraversal

import "testing"

//Test is a function to test postorderTraversal function

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	expected := []int{3, 2, 1}
	result := postorderTraversal(root)
	if len(result) != len(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v but got %v", expected, result)
			break
		}
	}

}
