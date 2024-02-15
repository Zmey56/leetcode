package hasPathSum

import "testing"

//Test for function hasPathSum

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  bool
	}{
		{"Testcase 1", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		}, 22, true},
		{"Testcase 2", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}, 1, false},
		{"Testcase 3", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}, 2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := hasPathSum(test.root, test.targetSum)
			if got != test.expected {
				t.Errorf("hasPathSum(%v, %v) = %v, want %v", test.root, test.targetSum, got, test.expected)
			}
		})
	}
}
