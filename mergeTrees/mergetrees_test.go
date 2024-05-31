package mergeTrees

import "testing"

// Test for the mergeTrees function

func Test_mergeTrees(t *testing.T) {
	// Test case where root1 and root2 are nil
	root1 := &TreeNode{1, nil, nil}
	root2 := &TreeNode{1, nil, nil}
	expected := &TreeNode{2, nil, nil}
	result := mergeTrees(root1, root2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where root1 is nil
	root1 = nil
	root2 = &TreeNode{1, nil, nil}
	expected = &TreeNode{1, nil, nil}
	result = mergeTrees(root1, root2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where root2 is nil
	root1 = &TreeNode{1, nil, nil}
	root2 = nil
	expected = &TreeNode{1, nil, nil}
	result = mergeTrees(root1, root2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where root1 and root2 have no common nodes
	root1 = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	root2 = &TreeNode{3, &TreeNode{4, nil, nil}, nil}
	expected = &TreeNode{4, &TreeNode{6, nil, nil}, nil}
	result = mergeTrees(root1, root2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where root1 and root2 have one common node
	root1 = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	root2 = &TreeNode{1, &TreeNode{4, nil, nil}, nil}
	expected = &TreeNode{2, &TreeNode{6, nil, nil}, nil}
	result = mergeTrees(root1, root2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

// Helper function to check if two slices are equal
func equal(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
