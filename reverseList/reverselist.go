package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example 1
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Example 1:")
	reversedHead1 := reverseList(head1)
	printList(reversedHead1)

	// Example 2
	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}

	fmt.Println("\nExample 2:")
	reversedHead2 := reverseList(head2)
	printList(reversedHead2)

	// Example 3
	var head3 *ListNode

	fmt.Println("\nExample 3:")
	reversedHead3 := reverseList(head3)
	printList(reversedHead3)
}
