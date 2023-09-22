package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	values := []int{}

	current := head

	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	maxTwinSum := 0

	left, right := 0, len(values)-1

	for left < right {
		twinSum := values[left] + values[right]
		if twinSum > maxTwinSum {
			maxTwinSum = twinSum
		}
		left++
		right--
	}
	return maxTwinSum

}

func main() {
	// Example 1
	head1 := &ListNode{Val: 5}
	head1.Next = &ListNode{Val: 4}
	head1.Next.Next = &ListNode{Val: 2}
	head1.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println("Example 1:", pairSum(head1))

	// Example 2
	head2 := &ListNode{Val: 4}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 2}
	head2.Next.Next.Next = &ListNode{Val: 3}
	fmt.Println("Example 2:", pairSum(head2))

	// Example 3
	head3 := &ListNode{Val: 1}
	head3.Next = &ListNode{Val: 100000}
	fmt.Println("Example 3:", pairSum(head3))
}
