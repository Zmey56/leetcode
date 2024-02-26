package hasCycle

import "testing"

//Test for hasCycle function

func TestHasCycle(t *testing.T) {
	tests := []struct {
		example   string
		input     []int
		exception bool
	}{
		{
			example:   "Test 1",
			input:     []int{3, 2, 0, -4},
			exception: true,
		},
		{
			example:   "Test 2",
			input:     []int{1, 2},
			exception: true,
		},
		{
			example:   "Test 3",
			input:     []int{1},
			exception: false,
		},
	}

	for _, test := range tests {
		head := createList(test.input)
		result := hasCycle(head)
		if result != test.exception {
			t.Errorf("Test %s failed, expected %v, got %v", test.example, test.exception, result)
		}
	}

}

// Create from slice a linked list
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}
