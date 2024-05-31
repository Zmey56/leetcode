package findTarget

//Given the root of a binary search tree and an integer k,
//return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return dfs(root, k, m)
}

func dfs(root *TreeNode, k int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}
	for key, _ := range m {
		if key+root.Val == k {
			return true
		}
	}

	m[root.Val] = struct{}{}

	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}
