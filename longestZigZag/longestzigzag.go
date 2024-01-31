package longestZigZag

//A ZigZag path for a binary tree is defined as follow:
//
//Choose any node in the binary tree and a direction (right or left).
//If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
//Change the direction from right to left or from left to right.
//Repeat the second and third steps until you can't move in the tree.
//Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
//
//Return the longest ZigZag path contained in that tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	dfs(root.Left, 1, true, &max)
	dfs(root.Right, 1, false, &max)
	return max
}

func dfs(root *TreeNode, length int, isLeft bool, max *int) {
	if root == nil {
		return
	}
	if length > *max {
		*max = length
	}
	if isLeft {
		dfs(root.Right, length+1, false, max)
		dfs(root.Left, 1, true, max)
	} else {
		dfs(root.Left, length+1, true, max)
		dfs(root.Right, 1, false, max)
	}
}
