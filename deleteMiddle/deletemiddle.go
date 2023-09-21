package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	arr := []int{}

	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
		arr = append(arr, current.Val)
	}

	if len(arr) == 0 {
		return nil
	} else if len(arr) == 1 {
		return head
	} else {
		target := length / 2

		dummy := &ListNode{Val: arr[0]}
		curDummy := dummy

		for i := 1; i < len(arr); i++ {
			if i != target {
				newNode := &ListNode{Val: arr[i]}
				curDummy.Next = newNode
				curDummy = newNode
			}

		}

		return dummy
	}
}
