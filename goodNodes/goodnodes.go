package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	return countGoodNodes(root, math.MinInt64)

}

func countGoodNodes(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count++
		maxSoFar = node.Val
	}

	count += countGoodNodes(node.Left, maxSoFar)
	count += countGoodNodes(node.Right, maxSoFar)

	return count
}

func main() {
	nodeOne := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Print(goodNodes(nodeOne))
}
