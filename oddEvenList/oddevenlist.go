package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyEven := &ListNode{Next: head.Next}
	odd, even := head, head.Next

	for even != nil && odd != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = dummyEven.Next

}
