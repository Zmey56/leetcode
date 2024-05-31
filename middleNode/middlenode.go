package middleNode

//Given the head of a singly linked list, return the middle node of the linked list.
//
//If there are two middle nodes, return the second middle node.

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	dummy := &ListNode{}
	tmpSlice := []int{}
	for head != nil {
		length++
		tmpSlice = append(tmpSlice, head.Val)
		head = head.Next
	}

	mid := length / 2
	tmpSlice = tmpSlice[mid:]

	current := dummy
	for i := 0; i < len(tmpSlice); i++ {
		current.Next = &ListNode{Val: tmpSlice[i]}
		current = current.Next
	}
	return dummy.Next
}
