package rightSideView

//Given the root of a binary tree, imagine yourself standing on the right side of it,
//return the values of the nodes you can see ordered from top to bottom.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if i == size-1 {
				res = append(res, cur.Val)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[size:]
	}
	return res
}
