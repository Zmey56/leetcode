package isPalindromeTwo

import "testing"

//Test for isPalindrome

func TestIsPalindrome(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	got := isPalindrome(head)
	if got != true {
		t.Errorf("Expected true but got %v", got)
	}
}
