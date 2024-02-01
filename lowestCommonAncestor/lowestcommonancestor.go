package lowestCommonAncestor

//Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
//
//According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as
//the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

//Complexity Analysis
//
//Time complexity : O(N)O(N), where NN is the number of nodes in the binary tree. In the worst case we might be visiting
//all the nodes of the binary tree.
//
//Space complexity : O(N)O(N). This is because the maximum amount of space utilized by the recursion stack would be NN
//since the height of a skewed binary tree could be NN.

//Compare this snippet from lowestCommonAncestor/lowestcommonancestor_test.go:
//package lowestCommonAncestor
//
//import "testing"
//
//func TestLowestCommonAncestor(t *testing.T) {
