package inorderTraversal

//Given the root of a binary tree, return the inorder traversal of its nodes' values.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	node := root

	for node != nil{
		if node.Left == nil
	}

}
