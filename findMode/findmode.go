package findMode

//Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.
//
//If the tree has more than one mode, return them in any order.
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than or equal to the node's key.
//The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
//Both the left and right subtrees must also be binary search trees.

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	f := make(freqMap)
	f.traverse(root)
	return f.mode()
}

type freqMap map[int]int

func (f freqMap) mode() []int {
	modes := make([]int, 0, len(f))
	var maxFreq int
	for elem, freq := range f {
		if freq > maxFreq {
			modes = modes[:0]
			maxFreq = freq
		}
		if freq == maxFreq {
			modes = append(modes, elem)
		}
	}
	return modes
}

func (f freqMap) traverse(node *TreeNode) {
	if node == nil {
		return
	}

	f[node.Val]++
	f.traverse(node.Left)
	f.traverse(node.Right)
}
