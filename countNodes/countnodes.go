package countNodes

//Given the root of a complete binary tree, return the number of the nodes in the tree.
//
//According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree,
//and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
//
//Design an algorithm that runs in less than O(n) time complexity.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}

	if root.Left != nil {
		count += countNodes(root.Left)
	}

	if root.Right != nil {
		count += countNodes(root.Right)
	}

	return count + 1
}
