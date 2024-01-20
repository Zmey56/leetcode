package swapPairs

import (
	"fmt"
	"testing"
)

// create test for function swapPairs
func TestSwapPairs(t *testing.T) {
	// Example 1
	head1 := &ListNode{Val: 5}
	head1.Next = &ListNode{Val: 4}
	head1.Next.Next = &ListNode{Val: 2}
	head1.Next.Next.Next = &ListNode{Val: 1}
	printList(swapPairs(head1))
	fmt.Println("--------------------")

	// Example 2
	head2 := &ListNode{Val: 4}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 2}
	head2.Next.Next.Next = &ListNode{Val: 3}
	printList(swapPairs(head2))
	fmt.Println("--------------------")

	// Example 3
	head3 := &ListNode{Val: 1}
	head3.Next = &ListNode{Val: 100000}
	printList(swapPairs(head3))
	fmt.Println("--------------------")

	//Example 4
	head4 := &ListNode{}
	printList(swapPairs(head4))
	fmt.Println("--------------------")

	//Example 5
	//Example 4
	head5 := &ListNode{Val: 1}
	printList(swapPairs(head5))
	fmt.Println("--------------------")

}
