package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

//Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying
//the values in the list's nodes (i.e., only nodes themselves may be changed.)

func swapPairs(head *ListNode) *ListNode {
	values := []int{}

	current := head

	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	// swap values
	for i := 0; i < len(values); i += 2 {
		if i+1 < len(values) {
			values[i], values[i+1] = values[i+1], values[i]
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

// function for printing linked list
func printList(head *ListNode) {
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
