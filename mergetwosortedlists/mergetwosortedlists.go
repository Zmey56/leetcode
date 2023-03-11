//21. Merge Two Sorted Lists
//Easy
//17.4K
//1.6K
//Companies
//You are given the heads of two sorted linked lists list1 and list2.
//
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
//Return the head of the merged linked list.
//
//
//
//Example 1:
//
//
//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//Example 2:
//
//Input: list1 = [], list2 = []
//Output: []
//Example 3:
//
//Input: list1 = [], list2 = [0]
//Output: [0]
//
//
//Constraints:
//
//The number of nodes in both lists is in the range [0, 50].
//-100 <= Node.val <= 100
//Both list1 and list2 are sorted in non-decreasing order.

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to hold the merged list
	dummy := &ListNode{}
	// Create a pointer to the last node of the merged list
	tail := dummy

	// Iterate over the two input lists, comparing and merging their nodes
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Append the remaining nodes from list1 or list2
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	// Return the merged list, skipping the dummy node
	return dummy.Next
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

	mergeTwoLists(l1, l2)
}
