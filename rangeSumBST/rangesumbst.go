package rangeSumBST

//Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes
//with a value in the inclusive range [low, high].

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum

}
