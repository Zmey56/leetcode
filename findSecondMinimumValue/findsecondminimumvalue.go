package findSecondMinimumValue

import "sort"

//Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree
//has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes.
//More formally, the property root.val = min(root.left.val, root.right.val) always holds.
//
//Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.
//
//If no such second minimum value exists, output -1 instead.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	valuesMap := make(map[int]struct{})

	valuesMap = findSecondMinimumValueHelper(root, valuesMap)

	values := make([]int, 0, len(valuesMap))
	for k := range valuesMap {
		values = append(values, k)
	}

	sort.Ints(values)

	if len(values) < 2 {
		return -1
	}

	return values[1]
}

func findSecondMinimumValueHelper(root *TreeNode, values map[int]struct{}) map[int]struct{} {
	if root == nil {
		return values
	}

	values[root.Val] = struct{}{}

	if root.Left != nil {
		values = findSecondMinimumValueHelper(root.Left, values)
	}

	if root.Right != nil {
		values = findSecondMinimumValueHelper(root.Right, values)
	}

	return values
}
