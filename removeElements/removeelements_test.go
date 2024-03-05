package removeElements

import "testing"

//Test for function removeElements

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    *ListNode
		val      int
		expected *ListNode
	}{
		{
			name:     "Test Case 1",
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}},
			val:      6,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
		{
			name:     "Test Case 2",
			input:    &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}},
			val:      7,
			expected: nil,
		},
		{
			name:     "Test Case 3",
			input:    &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}},
			val:      1,
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := removeElements(test.input, test.val)
			if !isListEqual(got, test.expected) {
				t.Errorf("Got %v, expected %v", got, test.expected)
			}
		})
	}
}

func isListEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
