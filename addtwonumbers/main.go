package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func addSliceToLinkedList(nums []int) *ListNode {
	// Create a dummy head node
	head := &ListNode{Val: 0, Next: nil}

	// Keep track of the current node
	current := head

	// Iterate over the slice and create a new node for each element
	for _, num := range nums {
		// Create a new node with the current element as its value
		newNode := &ListNode{Val: num, Next: nil}

		// Link the new node to the current node
		current.Next = newNode

		// Update the current node to be the new node
		current = newNode
	}

	// Return the head node (skipping the dummy head)
	return head.Next
}

func main() {
	l1s := []int{2, 4, 3}
	l2s := []int{5, 6, 4}

	l1 := addSliceToLinkedList(l1s)
	l2 := addSliceToLinkedList(l2s)

	addTwoNumbers(l1, l2)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy head node
	head := &ListNode{Val: 0, Next: nil}

	// Keep track of the current node and the carry
	current := head
	carry := 0

	// Loop through both linked lists
	for l1 != nil || l2 != nil {
		// Get the values of the current nodes, or 0 if the nodes are nil
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Calculate the sum and the carry
		log.Println("1 - VAL1", val1, "VAL2", val2, "CARRY", carry)
		sum := val1 + val2 + carry
		log.Println("SUM1", sum)
		carry = sum / 10
		sum = sum % 10
		log.Println("SUM2", sum)
		log.Println("2 - VAL1", val1, "VAL2", val2, "CARRY", carry)

		// Create a new node for the sum
		newNode := &ListNode{Val: sum, Next: nil}

		// Link the new node to the current node and update the current node
		current.Next = newNode
		current = newNode
	}

	// If there's still a carry, create a new node for it
	if carry > 0 {
		newNode := &ListNode{Val: carry, Next: nil}
		current.Next = newNode

		current = newNode
		log.Println(current)
	}

	log.Println("Result", head.Next)
	log.Println("Result", head)
	// Return the head node (skipping the dummy head)
	return head.Next
}
