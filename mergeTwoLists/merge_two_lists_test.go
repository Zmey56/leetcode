package mergeTwoLists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "Example 1",
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
		},
		{
			name:  "Example 2",
			list1: &ListNode{},
			list2: &ListNode{},
			want:  &ListNode{},
		},
		{
			name:  "Example 3",
			list1: &ListNode{Val: 0},
			list2: &ListNode{},
			want:  &ListNode{Val: 0},
		},
	}

	for _, tt := range tests {
		fmt.Println(tt.name)
		result := mergeTwoLists(tt.list1, tt.list2)
		if result.Val != tt.want.Val {
			t.Errorf("mergeTwoLists() = %v, want %v", result.Val, tt.want.Val)
		}
	}

}
