package removeElements

//Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val,
//and return the new head.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	return dummy.Next
}
