package maxDepthTwo

//Given a n-ary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//Nary-Tree input serialization is represented in their level order traversal, each group of children is separated
//by the null value (See examples).

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	queue := []*Node{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		depth++
	}

	return depth
}
