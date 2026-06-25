package middlelinkedlist

// 876. Middle of the Linked List

// Given the head of a singly linked list, return the middle node of the linked list.

// If there are two middle nodes, return the second middle node.

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    current := head
    count := 0
    for current != nil{
        count++
        current = current.Next
    }

    mid := count/2

    for i:= 0; i < mid; i++{
        if head == nil{
            return nil
        }
        head = head.Next
    }

    return head
}

func middleNodeV2(head *ListNode) *ListNode {
    slow := head.Next
    fast := head.Next.Next
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}
