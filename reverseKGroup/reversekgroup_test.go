package reverseKGroup

import (
	"log"
	"testing"
)

//create test for function reverseKGroup

func TestReverseKGroup(t *testing.T) {
	// Example 1
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(reverseKGroup(head1, 2))

	// Example 2
	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	head2.Next.Next.Next = &ListNode{Val: 4}
	head2.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(reverseKGroup(head2, 3))
}

func printList(head *ListNode) {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		println(head.Val)
		head = head.Next
	}
	log.Println(result)
}
