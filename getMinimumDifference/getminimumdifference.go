package getMinimumDifference

//Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var result int
	var prev *TreeNode
	result = 1<<63 - 1
	prev = nil
	inOrder(root, &prev, &result)
	return result
}

func inOrder(root *TreeNode, prev **TreeNode, result *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, prev, result)
	if *prev != nil {
		if root.Val-(*prev).Val < *result {
			*result = root.Val - (*prev).Val
		}
	}
	*prev = root
	inOrder(root.Right, prev, result)
}
