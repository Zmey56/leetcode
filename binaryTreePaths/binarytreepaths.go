package binaryTreePaths

import "strconv"

//Given the root of a binary tree, return all root-to-leaf paths in any order.
//
//A leaf is a node with no children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}
	dfs(root, strconv.Itoa(root.Val), &result)
	return result
}

func dfs(node *TreeNode, path string, result *[]string) {
	if node.Left == nil && node.Right == nil {
		*result = append(*result, path)
	}
	if node.Left != nil {
		dfs(node.Left, path+"->"+strconv.Itoa(node.Left.Val), result)
	}
	if node.Right != nil {
		dfs(node.Right, path+"->"+strconv.Itoa(node.Right.Val), result)
	}
}
