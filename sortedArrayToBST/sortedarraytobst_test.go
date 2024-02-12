package sortedArrayToBST

import "testing"

//TestSortedArrayToBST is a test function

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums   []int
		result *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}},
		{[]int{1, 3}, &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
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
		}}}

	for _, test := range tests {
		if result := sortedArrayToBST(test.nums); !isSameTree(result, test.result) {
			t.Errorf("sortedArrayToBST(%v) = %v, want %v", test.nums, result, test.result)
		}
	}
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
