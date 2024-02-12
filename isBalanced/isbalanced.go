package isBalanced

//Given a binary tree, determine if it is height-balanced.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1

}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
