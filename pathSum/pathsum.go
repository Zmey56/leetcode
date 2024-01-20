package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// Start the DFS traversal from the root
	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val == targetSum {
		count++
	}

	count += dfs(node.Left, targetSum-node.Val)
	count += dfs(node.Right, targetSum-node.Val)

	return count
}

func main() {
	// Example usage
	root1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}},
		},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}},
	}

	targetSum1 := 8
	result1 := pathSum(root1, targetSum1)
	fmt.Println("Example 1:", result1) // Output: 3

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
		},
	}

	targetSum2 := 22
	result2 := pathSum(root2, targetSum2)
	fmt.Println("Example 2:", result2) // Output: 3
}
