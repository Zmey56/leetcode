package postorder

// Given the root of an n-ary tree, return the postorder traversal of its nodes' values.
//
// Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}

	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)

	return result
}
