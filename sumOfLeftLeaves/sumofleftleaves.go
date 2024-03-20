package sumOfLeftLeaves

//Given the root of a binary tree, return the sum of all left leaves.
//
//A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

// TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SumOfLeftLeaves returns the sum of all left leaves
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
