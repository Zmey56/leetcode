package findTilt

//Given the root of a binary tree, return the sum of every tree node's tilt.
//
//The tilt of a tree node is the absolute difference between the sum of all left subtree node values and
//all right subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0.
//The rule is similar if the node does not have a right child.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var tilt int
	traverse(root, &tilt)
	return tilt
}

func traverse(node *TreeNode, tilt *int) int {
	if node == nil {
		return 0
	}

	left := traverse(node.Left, tilt)
	right := traverse(node.Right, tilt)

	*tilt += abs(left - right)

	return node.Val + left + right
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
