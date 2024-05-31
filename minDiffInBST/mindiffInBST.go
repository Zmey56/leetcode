package minDiffInBST

//Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var (
		prev    *TreeNode
		minDiff = 1<<63 - 1
	)
	inorder(root, &prev, &minDiff)
	return minDiff
}

func inorder(root *TreeNode, prev **TreeNode, minDiff *int) {
	if root == nil {
		return
	}
	inorder(root.Left, prev, minDiff)
	if *prev != nil {
		*minDiff = min(*minDiff, root.Val-(*prev).Val)
	}
	*prev = root
	inorder(root.Right, prev, minDiff)
}
