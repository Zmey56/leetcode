package minDepth

import "testing"

//Test for function minDepth

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name  string
		root  *TreeNode
		count int
	}{
		{"Testcase 1", &TreeNode{
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
		}, 2},
		{"Testcase 2", &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minDepth(test.root)
			if got != test.count {
				t.Errorf("minDepth(%v) = %v, want %v", test.root, got, test.count)
			}
		})
	}
}
