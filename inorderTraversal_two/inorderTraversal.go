package main

// 94. Binary Tree Inorder Traversal

// Given the root of a binary tree, return the inorder traversal of its nodes' values.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var result []int

    var traverse func(node *TreeNode)
    traverse = func(node *TreeNode){
        if node == nil{
            return 
        }

       	traverse(node.Left)...)
	result = append(result, node.Val)
	result = append(result, inorderTraversal(root.Right)...)
    }
    

	return result
}
