package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafSeq1 := []int{}
	leafSeq2 := []int{}

	var dfs func(root *TreeNode, sequence *[]int)
	dfs = func(root *TreeNode, sequence *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*sequence = append(*sequence, root.Val)
		}

		dfs(root.Left, sequence)
		dfs(root.Right, sequence)
	}

	dfs(root1, &leafSeq1)
	dfs(root2, &leafSeq2)

	if len(leafSeq1) != len(leafSeq2) {
		return false
	}

	for i, num := range leafSeq1 {
		if num != leafSeq2[i] {
			return false
		}
	}

	return true
}

func main() {
	// Create the two binary trees
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	// Check if the two trees are leaf-similar using BFS
	result := leafSimilar(root1, root2)
	fmt.Println(result) // Output: true
}
