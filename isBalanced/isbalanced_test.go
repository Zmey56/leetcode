package isBalanced

import "testing"

//TetIsBalanced is a test function

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		result bool
	}{
		{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}, true},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}, false},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}, true},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}, true},
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 2,
			}}, false},
	}
	for _, test := range tests {
		if result := isBalanced(test.root); result != test.result {
			t.Errorf("isBalanced(%v) = %v, want %v", test.root, result, test.result)
		}
	}
}
