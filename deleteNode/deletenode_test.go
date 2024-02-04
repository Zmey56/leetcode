package deleteNode

import "testing"

// Test for function deleteNode
func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		key  int
		want *TreeNode
	}{
		{
			name: "Test 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
			key: 3,
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "Test 2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
			key: 0,
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   6,
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}
	for _, test := range tests {
		got := deleteNode(test.root, test.key)
		if !isSameTree(got, test.want) {
			t.Errorf("deleteNode(%v, %v) = %v; want %v", test.root, test.key, got, test.want)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
