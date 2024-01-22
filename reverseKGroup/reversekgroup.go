package reverseKGroup

//Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not
//a multiple of k then left-out nodes, in the end, should remain as it is.
//
//You may not alter the values in the list's nodes, only nodes themselves may be changed.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	values := []int{}

	current := head

	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	// reverse values
	for i := 0; i < len(values); i += k {
		if i+k <= len(values) {
			for j := 0; j < k/2; j++ {
				values[i+j], values[i+k-1-j] = values[i+k-1-j], values[i+j]
			}
		}
	}

	// create new list
	newHead := &ListNode{}
	current = newHead

	//add values to new list
	for i := 0; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next

	}

	// return new list
	return newHead.Next
}
