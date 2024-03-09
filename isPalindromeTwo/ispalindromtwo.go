package isPalindromeTwo

//Given the head of a singly linked list, return true if it is a
//palindrome or false otherwise.

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	result := []int{}

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-1-i] {
			return false
		}
	}

	return true

}
