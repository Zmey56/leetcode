package inorderTraversal

//Given the root of a binary tree, return the inorder traversal of its nodes' values.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var values []int
	values = append(values, inorderTraversal(root.Left)...)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)

	return values
}
