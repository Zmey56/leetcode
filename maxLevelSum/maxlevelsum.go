package maxLevelSum

import "math"

// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
//
//Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	max := math.MinInt32
	level := 1
	res := 1

	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		tmp := queue
		queue = []*TreeNode{}
		sum := 0

		for _, node := range tmp {
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > max {
			max = sum
			res = level
		}

		level++
	}

	return res
}
