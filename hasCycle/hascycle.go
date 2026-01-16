package hasCycle

//Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
//There is a cycle in a linked list if there is some node in the list that can be reached again by continuously
//following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
//Note that pos is not passed as a parameter.
//
//Return true if there is a cycle in the linked list. Otherwise, return false.

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCycleWithMap(head *ListNode) bool {
	if head == nil {
		return false
	}
	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}

func hasCycleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
