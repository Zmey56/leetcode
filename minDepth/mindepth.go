package minDepth

import "log"

//Given a binary tree, find its minimum depth.
//
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
//Note: A leaf is a node with no children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	rightSlice := []int{}
	leftSlice := []int{}
	right, rightSlice := minDepthRight(root, rightSlice)
	left, leftSlice := minDepthLeft(root, leftSlice)
	log.Println("right", right, "left", left)
	log.Println("rightSlice", rightSlice, "leftSlice", leftSlice)
	return min(len(rightSlice), len(leftSlice))
}

func minDepthRight(root *TreeNode, rightSlice []int) (int, []int) {
	rightSlice = append(rightSlice, root.Val)
	if root == nil {
		return 0, rightSlice
	}
	if root.Left == nil && root.Right == nil {
		return 1, rightSlice
	}
	return minDepthRight(root.Right, rightSlice)
}

func minDepthLeft(root *TreeNode, leftSlice []int) (int, []int) {
	leftSlice = append(leftSlice, root.Val)
	if root == nil {
		return 0, leftSlice
	}
	if root.Left == nil && root.Right == nil {
		return 1, leftSlice
	}
	return minDepthLeft(root.Left, leftSlice)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
