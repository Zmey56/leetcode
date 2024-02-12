package sortedArrayToBST

//Given an integer array nums where the elements are sorted in ascending order, convert it to a
//height-balanced binary search tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(numc []int) *TreeNode {
	if len(numc) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   numc[len(numc)/2],
		Left:  sortedArrayToBST(numc[:len(numc)/2]),
		Right: sortedArrayToBST(numc[len(numc)/2+1:]),
	}

}
