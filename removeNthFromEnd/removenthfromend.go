package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy

	// Move the first pointer n+1 steps ahead.
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	// Move the first pointer to the end, maintaining the gap between first and second.
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Skip the nth node from the end.
	second.Next = second.Next.Next

	return dummy.Next
}

func createLinkedList(values []int) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	current := dummy
	for _, val := range values {
		current.Next = &ListNode{Val: val, Next: nil}
		current = current.Next
	}
	return dummy.Next
}

func displayLinkedList(head *ListNode) []int {
	result := []int{}
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func main() {
	fmt.Println("Example 1:")
	head1 := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original List: ")
	displayLinkedList(head1)
	n1 := 2
	newHead1 := removeNthFromEnd(head1, n1)
	fmt.Print("List after removing ", n1, "th node from the end: ")
	displayLinkedList(newHead1)

	fmt.Println("\nExample 2:")
	head2 := createLinkedList([]int{1})
	fmt.Print("Original List: ")
	displayLinkedList(head2)
	n2 := 1
	newHead2 := removeNthFromEnd(head2, n2)
	fmt.Print("List after removing ", n2, "th node from the end: ")
	displayLinkedList(newHead2)

	fmt.Println("\nExample 3:")
	head3 := createLinkedList([]int{1, 2})
	fmt.Print("Original List: ")
	displayLinkedList(head3)
	n3 := 1
	newHead3 := removeNthFromEnd(head3, n3)
	fmt.Print("List after removing ", n3, "th node from the end: ")
	displayLinkedList(newHead3)
}
