package isSubtree

//Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with
//the same structure and node values of subRoot and false otherwise.
//
//A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
//The tree tree could also be considered as a subtree of itself.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
