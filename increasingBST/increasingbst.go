package increasingBST

// Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now
//the root of the tree, and every node has no left child and only one right child.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	//var stack []*TreeNode
	var inorder func(node *TreeNode)
	var dummyNode = &TreeNode{}
	var curr = dummyNode

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		curr.Right = node
		node.Left = nil
		curr = node
		inorder(node.Right)
	}

	inorder(root)
	return dummyNode.Right

}
