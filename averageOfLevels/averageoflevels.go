package averageOfLevels

//Given the root of a binary tree, return the average value of the nodes on each level in the form of an array.
//Answers within 10-5 of the actual answer will be accepted.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		count := 0
		var next []*TreeNode
		for _, node := range queue {
			sum += node.Val
			count++
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(count))
		queue = next
	}
	return result

}
