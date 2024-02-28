package getIntersectionNode

//Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
//If the two linked lists have no intersection at all, return null.
//
//For example, the following two linked lists begin to intersect at node c1:
//
//
//The test cases are generated such that there are no cycles anywhere in the entire linked structure.

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// get the length of the two linked lists
	lenA := getLength(headA)
	lenB := getLength(headB)

	// move the pointer of the longer linked list to the same length as the shorter linked list
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	// move the pointer of the two linked lists at the same time until they meet
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

// get the length of the linked list
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
