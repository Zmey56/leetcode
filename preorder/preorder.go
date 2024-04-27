package preorder

//Given the root of an n-ary tree, return the preorder traversal of its nodes' values.
//
//Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)\

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}

	return result

}
